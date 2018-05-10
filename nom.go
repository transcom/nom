package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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

		ssn := record[fields["EMPLID"]]
		orderCntlNbr := record[fields["N_ORDER_CNTL_NBR"]]
		// TODO construct new orders num by converting SSN to EDIPI, then combining that with the order_cntl_nbr expanded into the julian date concatenated with the 4 digit year
		ordersDate := record[fields["N_ORD_DT"]]
		modNbr, modNbrErr := strconv.Atoi(record[fields["N_MOD_NBR"]])
		if modNbrErr != nil {
			modNbr = 0
		}
		modNum, modNumErr := strconv.Atoi(record[fields["N_MOD_NUM"]])
		if modNumErr != nil {
			modNum = 0
		}
		seqNum := modNbr + modNum

		var status string
		if record[fields["N_OBLG_STATUS"]] == "D" {
			status = "canceled"
		} else {
			status = "authorized"
		}
		// TODO understand TDY en route with the N_OBLG_LEG_NBR field
		rateRank := record[fields["N_RATE_RANK"]]
		rank := RankFromAbbreviation(rateRank)

		purpose := record[fields["N_CIC_PURP"]]
		var ordersType swagger.OrdersType
		if rank.officer {
			ordersType = officerOrdersTypes[purpose]
		} else {
			ordersType = enlistedOrdersTypes[purpose]
		}

		// The sailor's name, in the format LASTNAME,FIRSTNAME (optional MI) (optional suffix)
		name := record[fields["NAME"]]

		losingUnitName := record[fields["N_DET_HPORT"]]
		losingUnitIdentCode := fmt.Sprintf("N%05s", record[fields["N_UIC_DETACH"]])
		losingUnitCity := record[fields["N_PDS_CITY"]]
		losingUnitState := record[fields["N_PDS_STATE"]]
		losingUnitCountry := record[fields["N_PDS_CNTRY"]]

		estArrivalDate := record[fields["N_EST_ARRIVAL_DT"]]

		gainingUnitName := record[fields["N_ULT_HPORT"]]
		gainingUnitIdentCode := fmt.Sprintf("N%05s", record[fields["N_UIC_ULT_DTY_STA"]])
		gainingUnitCity := record[fields["N_ULT_CITY"]]
		gainingUnitState := record[fields["N_ULT_STATE"]]
		gainingUnitCountry := record[fields["N_ULT_CNTRY"]]

		/*
			| N_NON_ENT_IND | If 'Y', then this is a 'Cost Order' with obligated moving expenses. If 'N', then this is a 'No Cost Order', i.e., a PCA w/o PCS (Permanent Change of Assignment without Permanent Change of Station), and has no moving expenses. |
		*/

		var dependents bool
		if record[fields["N_NUM_DEPN"]] == "Y" {
			dependents = true
		} else {
			dependents = false
		}

		tacSdn := record[fields["TAC_SDN"]]
		tac := tacSdn[len(tacSdn)-4:]
		sdn := tacSdn[:len(tacSdn)-4]

		fmt.Printf("%s %s (%d):\n", ssn, orderCntlNbr, seqNum)
		fmt.Printf("  %s %s %s (%s)\n", rateRank, rank.title, name, strings.ToUpper((string(rank.paygrade))))
		fmt.Printf("  Has Dependents: %t\n", dependents)
		fmt.Println("  ordersDate: " + ordersDate)
		fmt.Println("  status: " + status)
		fmt.Println("  ordersType: " + ordersType)
		fmt.Println("  estArrivalDate: " + estArrivalDate)
		fmt.Printf("  Losing Unit: %s (%s) %s, %s %s\n", losingUnitName, losingUnitIdentCode, losingUnitCity, losingUnitState, losingUnitCountry)
		fmt.Printf("  Gaining Unit: %s (%s) %s, %s %s\n", gainingUnitName, gainingUnitIdentCode, gainingUnitCity, gainingUnitState, gainingUnitCountry)
		fmt.Println("  SDN: " + sdn)
		fmt.Println("  TAC: " + tac)
	}
}
