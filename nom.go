package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/transcom/nom/pkg/swagger"
)

var enlistedOrdersTypes = map[string]swagger.OrdersType{
	"1": swagger.IPCOT,           // IPCOT In-place consecutive overseas travel
	"8": swagger.OTEIP,           // Overseas Tour Extension Incentive Program (OTEIP)
	"9": swagger.TRAINING,        // NAVCAD (Naval Cadet) Training
	"A": swagger.ACCESSION,       // Accession Travel Recruits
	"B": swagger.ACCESSION,       // Non-recruit Accession Travel
	"C": swagger.TRAINING,        // Training Travel
	"D": swagger.OPERATIONAL,     // Operational Travel
	"E": swagger.SEPARATION,      // Separation Travel
	"F": swagger.UNIT_MOVE,       // Organized Unit/Homeport Change
	"G": swagger.ACCESSION,       // Midshipman Accession Travel
	"H": swagger.SPECIAL_PURPOSE, // Special Purpose Reimbursable
	"I": swagger.ACCESSION,       // NAVCAD(Naval Cadet) Accession
	"J": swagger.ACCESSION,       // Accession Travel Recruits
	"K": swagger.ACCESSION,       // Non-recruit Accession Travel
	"L": swagger.TRAINING,        // Training Travel
	"M": swagger.ROTATIONAL,      // Rotational Travel
	"N": swagger.SEPARATION,      // Separation Travel
	"O": swagger.UNIT_MOVE,       // Organized Unit/Homeport Change
	"P": swagger.SEPARATION,      // Midshipman Separation Travel
	"R": swagger.OPERATIONAL,     // Misc. Operational Non-member
	"X": swagger.EMERGENCY_EVAC,  // EMERGENCY NON-MEMBER EVACS
	"Y": swagger.ROTATIONAL,      // Misc. Rotational Non-member
	"Z": swagger.SEPARATION,      // NAVCAD(Naval Cadet) Separation
}

var officerOrdersTypes = map[string]swagger.OrdersType{
	"0": swagger.IPCOT,           // IPCOT In-place consecutive overseas travel
	"2": swagger.ACCESSION,       // Accession Travel
	"3": swagger.TRAINING,        // Training Travel
	"4": swagger.OPERATIONAL,     // Operational Travel
	"5": swagger.SEPARATION,      // Separation Travel
	"6": swagger.UNIT_MOVE,       // Organized Unit/Homeport Change
	"7": swagger.EMERGENCY_EVAC,  // Emergency Non-member Evac
	"H": swagger.SPECIAL_PURPOSE, // Special Purpose Reimbursable
	"Q": swagger.ROTATIONAL,      // Misc. Rotational Non-member
	"S": swagger.ACCESSION,       // Accession Travel
	"T": swagger.TRAINING,        // Training Travel
	"U": swagger.ROTATIONAL,      // Rotational Travel
	"V": swagger.SEPARATION,      // Separation Travel
	"W": swagger.UNIT_MOVE,       // Organized Unit/Homeport Change
	"X": swagger.ROTATIONAL,      // Misc. Rotational Non-member
}

func main() {
	inputPath := os.Args[1]
	fileReader, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	csvReader := csv.NewReader(fileReader)

	// First line contains the column headers; make a hash table that keys on the header with the column index as the value
	headers, err := csvReader.Read()
	if err != nil {
		if err == io.EOF {
			log.Fatal("Empty file; no headers found")
		} else {
			log.Fatal(err)
		}
		os.Exit(1)
	}
	fields := make(map[string]int)
	for i := 0; i < len(headers); i++ {
		fields[headers[i]] = i
	}

	// every subsequent line can now be picked apart using this information
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
			os.Exit(1)
		}

		var rev swagger.Revision
		affiliation := swagger.NAVY
		rev.Member = new(swagger.Member)
		rev.Member.Affiliation = &affiliation
		// The sailor's name, in the format LASTNAME,FIRSTNAME (optional MI) (optional suffix)
		// TODO - when we get appropriately formatted files, parse the name into its components
		rev.Member.FamilyName = record[fields["NAME"]]

		//ssn := record[fields["EMPLID"]]
		//orderCntlNbr := record[fields["N_ORDER_CNTL_NBR"]]
		// TODO construct new orders num by converting SSN to EDIPI and converting the N_ORDER_CNTL_NBR to
		// ISO 8601 date, combining the two as {iso8601date}-{edipi}

		const usaDate = "1/2/06"
		d, _ := time.Parse(usaDate, record[fields["N_ORD_DT"]])
		rev.DateIssued = d

		modNbr, modNbrErr := strconv.Atoi(record[fields["N_MOD_NBR"]])
		if modNbrErr != nil {
			modNbr = 0
		}
		modNum, modNumErr := strconv.Atoi(record[fields["N_MOD_NUM"]])
		if modNumErr != nil {
			modNum = 0
		}
		rev.SeqNum = int32(modNbr + modNum)

		if record[fields["N_OBLG_STATUS"]] == "D" {
			rev.Status = "canceled"
		} else {
			rev.Status = "authorized"
		}
		rateRank := record[fields["N_RATE_RANK"]]
		rank := RankFromAbbreviation(rateRank)
		paygrade := rank.paygrade
		rev.Member.Title = rank.title
		rev.Member.Rank = &paygrade

		purpose := record[fields["N_CIC_PURP"]]
		var ordersType swagger.OrdersType
		if rank.officer {
			ordersType = officerOrdersTypes[purpose]
		} else {
			ordersType = enlistedOrdersTypes[purpose]
		}
		rev.OrdersType = &ordersType

		rev.LosingUnit = new(swagger.Unit)
		rev.LosingUnit.Name = record[fields["N_DET_HPORT"]]
		rev.LosingUnit.Uic = fmt.Sprintf("N%05s", record[fields["N_UIC_DETACH"]])
		rev.LosingUnit.City = record[fields["N_PDS_CITY"]]
		rev.LosingUnit.Locality = record[fields["N_PDS_STATE"]]
		rev.LosingUnit.Country = record[fields["N_PDS_CNTRY"]]

		d, _ = time.Parse(usaDate, record[fields["N_EST_ARRIVAL_DT"]])
		year, month, day := d.Date()
		rev.ReportNoLaterThan = fmt.Sprintf("%d-%02d-%02d", year, month, day)

		rev.GainingUnit = new(swagger.Unit)
		rev.GainingUnit.Name = record[fields["N_ULT_HPORT"]]
		rev.GainingUnit.Uic = fmt.Sprintf("N%05s", record[fields["N_UIC_ULT_DTY_STA"]])
		rev.GainingUnit.City = record[fields["N_ULT_CITY"]]
		rev.GainingUnit.Locality = record[fields["N_ULT_STATE"]]
		rev.GainingUnit.Country = record[fields["N_ULT_CNTRY"]]

		if record[fields["N_NON_ENT_IND"]] == "Y" {
			rev.PcaWithoutPcs = true
		} else {
			rev.PcaWithoutPcs = false
		}

		if record[fields["N_NUM_DEPN"]] == "Y" {
			rev.HasDependents = true
		} else {
			rev.HasDependents = false
		}

		tacSdn := record[fields["TAC_SDN"]]
		rev.PcsAccounting = new(swagger.Accounting)
		rev.PcsAccounting.Tac = tacSdn[len(tacSdn)-4:]
		rev.PcsAccounting.Sdn = tacSdn[:len(tacSdn)-4]

		bodyBuf := &bytes.Buffer{}
		encoder := json.NewEncoder(bodyBuf)
		encoder.SetIndent("", "  ")
		encoder.Encode(rev)

		fmt.Print(bodyBuf.String())
	}
}
