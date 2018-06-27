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
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/transcom/nom/pkg/gen/ordersapi/models"
)

var suffixes = []string{"JR", "SR", "II", "III", "IV", "V"}

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

		var rev models.Revision
		rev.Member = new(models.Member)
		rev.Member.Affiliation = models.AffiliationNavy
		// The sailor's name, in the format LASTNAME,FIRSTNAME (optional MI) (optional suffix)
		fullname := record[fields["Service Member Name"]]
		names := strings.SplitN(fullname, ",", 2)
		rev.Member.FamilyName = new(string)
		*rev.Member.FamilyName = names[0]
		names = strings.Fields(names[1])
		rev.Member.GivenName = new(string)
		*rev.Member.GivenName = names[0]
		if len(names) > 1 {
			if stringInSlice(names[len(names)-1], suffixes) {
				rev.Member.Suffix = names[len(names)-1]
				if len(names) > 2 {
					rev.Member.MiddleName = strings.Join(names[1:len(names)-1], " ")
				}
			} else {
				rev.Member.MiddleName = strings.Join(names[1:], " ")
			}
		}

		daysStarting31Dec1899, daysError := strconv.Atoi(record[fields["Order Create/Modification Date"]])
		dateIssued := time.Date(1899, time.December, 30+daysStarting31Dec1899, 0, 0, 0, 0, time.Local)
		rev.DateIssued = strfmt.DateTime(dateIssued)

		orderModNbr, orderModNbrErr := strconv.Atoi(record[fields["Order Modification Number"]])
		if orderModNbrErr != nil {
			orderModNbr = 0
		}
		obligModNbr, obligModNbrErr := strconv.Atoi(record[fields["Obligation Modification Number"]])
		if obligModNbrErr != nil {
			obligModNbr = 0
		}
		rev.SeqNum = new(int64)
		*rev.SeqNum = int64(orderModNbr + obligModNbr)

		if record[fields["Obligation Status Code"]] == "D" {
			rev.Status = models.RevisionStatusCanceled
		} else {
			rev.Status = models.RevisionStatusAuthorized
		}
		rev.Member.Title = record[fields["Rank Classification  Description"]]
		categorizedRank := paygradeToRank[record[fields["Paygrade"]]]
		rev.Member.Rank = categorizedRank.paygrade

		purpose := record[fields["CIC Purpose Information Code (OBLGTN)"]]
		if categorizedRank.officer {
			rev.OrdersType = officerOrdersTypes[purpose]
		} else {
			rev.OrdersType = enlistedOrdersTypes[purpose]
		}

		rev.LosingUnit = new(models.Unit)
		if name := strings.TrimSpace(record[fields["Detach UIC Home Port"]]); len(name) > 0 {
			rev.LosingUnit.Name = name
		}
		if uic := strings.TrimSpace(record[fields["Detach UIC"]]); len(uic) > 0 {
			rev.LosingUnit.Uic = fmt.Sprintf("N%05s", uic)
		}
		if city := strings.TrimSpace(record[fields["Detach UIC City Name"]]); len(city) > 0 {
			rev.LosingUnit.City = city
		}
		if state := strings.TrimSpace(record[fields["Detach State Code"]]); len(state) > 0 {
			rev.LosingUnit.Locality = state
		}
		if country := strings.TrimSpace(record[fields["Detach Country Code"]]); len(country) > 0 {
			rev.LosingUnit.Country = country
		}

		daysStarting31Dec1899, daysError = strconv.Atoi(record[fields["Ultimate Estimated Arrival Date"]])
		if daysError == nil {
			estArrivalDate := time.Date(1899, time.December, 30+daysStarting31Dec1899, 0, 0, 0, 0, time.Local)
			rev.ReportNoLaterThan = new(strfmt.Date)
			*rev.ReportNoLaterThan = strfmt.Date(estArrivalDate)
		}

		rev.GainingUnit = new(models.Unit)
		if name := strings.TrimSpace(record[fields["Ultimate UIC Home Port"]]); len(name) > 0 {
			rev.GainingUnit.Name = name
		}
		if uic := strings.TrimSpace(record[fields["Ultimate UIC"]]); len(uic) > 0 {
			rev.GainingUnit.Uic = fmt.Sprintf("N%05s", uic)
		}
		if city := strings.TrimSpace(record[fields["Ultimate UIC City Name"]]); len(city) > 0 {
			rev.GainingUnit.City = city
		}
		if state := strings.TrimSpace(record[fields["Ultimate State Code"]]); len(state) > 0 {
			rev.GainingUnit.Locality = state
		}
		if country := strings.TrimSpace(record[fields["Ultimate Country Code"]]); len(country) > 0 {
			rev.GainingUnit.Country = country
		}

		if record[fields["Entitlement Indicator"]] == "Y" {
			rev.NoCostMove = false
		} else {
			rev.NoCostMove = true
		}

		rev.HasDependents = new(bool)
		*rev.HasDependents = record[fields["Count of Dependents Participating in Move (STATIC)"]] != "0"

		if tdyEnRoute, tdyError := strconv.Atoi(record[fields["Count of Intermediate Stops (STATIC)"]]); tdyError == nil {
			rev.TdyEnRoute = tdyEnRoute > 0
		}

		rev.PcsAccounting = new(models.Accounting)
		rev.PcsAccounting.Tac = record[fields["TAC"]]

		bodyBuf := &bytes.Buffer{}
		encoder := json.NewEncoder(bodyBuf)
		encoder.SetIndent("", "  ")
		encoder.Encode(rev)

		fmt.Print(bodyBuf.String())
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
