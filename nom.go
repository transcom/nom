package main

import (
	"bytes"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	runtimeClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/namsral/flag"
	"github.com/tcnksm/go-input"
	"pault.ag/go/pksigner"

	"github.com/transcom/nom/pkg/gen/ordersapi/client"
	"github.com/transcom/nom/pkg/gen/ordersapi/client/operations"
	"github.com/transcom/nom/pkg/gen/ordersapi/models"
)

var suffixes = []string{"JR", "SR", "II", "III", "IV", "V"}

func main() {
	var (
		pkcs11ModulePath string
		tokenLabel       string
		certLabel        string
		keyLabel         string
		keyPath          string
		certPath         string
		host             string
		port             uint
		insecure         bool
	)
	flag.StringVar(&pkcs11ModulePath, "pkcs11module", "", "Smart card: Path to the PKCS11 module to use")
	flag.StringVar(&tokenLabel, "tokenlabel", "", "Smart card: name of the token to use")
	flag.StringVar(&certLabel, "certlabel", "Certificate for PIV Authentication", "Smart card: label of the public cert")
	flag.StringVar(&keyLabel, "keylabel", "PIV AUTH key", "Smart card: label of the private key")
	flag.StringVar(&keyPath, "key", "", "Certificate from file: Path to the client certificate's private key")
	flag.StringVar(&certPath, "cert", "", "Certificate from file: Path to the client TLS Certificate")
	flag.StringVar(&host, "host", "orders.move.mil", "Host name to send the orders to")
	flag.UintVar(&port, "port", 443, "Remote port number to connect to")
	flag.BoolVar(&insecure, "insecure", false, "Skip TLS verification and validation")

	flag.Parse()

	var httpClient *http.Client

	// The client certificate comes either from a file OR from a smart card
	if pkcs11ModulePath != "" {
		pkcsConfig := pksigner.Config{
			Module:           pkcs11ModulePath,
			CertificateLabel: certLabel,
			PrivateKeyLabel:  keyLabel,
			TokenLabel:       tokenLabel,
		}

		store, err := pksigner.New(pkcsConfig)
		if err != nil {
			log.Fatal(err)
		}
		defer store.Close()

		inputUI := &input.UI{
			Writer: os.Stdout,
			Reader: os.Stdin,
		}

		pin, err := inputUI.Ask("PIN", &input.Options{
			Default:     "",
			HideOrder:   true,
			HideDefault: true,
			Required:    true,
			Loop:        true,
			Mask:        true,
			ValidateFunc: func(input string) error {
				matched, matchErr := regexp.Match("^\\d+$", []byte(input))
				if matchErr != nil {
					return matchErr
				}
				if !matched {
					return errors.New("Invalid")
				}
				return nil
			},
		})
		if err != nil {
			os.Exit(1)
		}

		err = store.Login(pin)
		if err != nil {
			log.Fatal(err)
		}

		cert, err := store.TLSCertificate()
		if err != nil {
			panic(err)
		}
		// #nosec b/c gosec triggers on InsecureSkipVerify
		tlsConfig := &tls.Config{
			Certificates:       []tls.Certificate{*cert},
			InsecureSkipVerify: insecure,
			MinVersion:         tls.VersionTLS12,
			MaxVersion:         tls.VersionTLS12,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{
			TLSClientConfig: tlsConfig,
		}
		httpClient = &http.Client{
			Transport: transport,
		}
	} else {
		var err error
		httpClient, err = runtimeClient.TLSClient(runtimeClient.TLSClientOptions{Key: keyPath, Certificate: certPath, InsecureSkipVerify: insecure})
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	inputPath := flag.Arg(0)
	fileReader, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(fileReader)

	hostWithPort := fmt.Sprintf("%s:%d", host, port)
	myRuntime := runtimeClient.NewWithClient(hostWithPort, client.DefaultBasePath, []string{"https"}, httpClient)
	myRuntime.EnableConnectionReuse()
	myRuntime.SetDebug(true)

	// First line contains the column headers; make a hash table that keys on the header with the column index as the value
	headers, err := csvReader.Read()
	if err != nil {
		if err == io.EOF {
			log.Fatal("Empty file; no headers found")
		} else {
			log.Fatal(err)
		}
	}
	fields := make(map[string]int)
	for i := 0; i < len(headers); i++ {
		fields[headers[i]] = i
	}

	ordersGateway := client.New(myRuntime, nil)

	// every subsequent line can now be picked apart using this information
	for record, recordErr := csvReader.Read(); recordErr == nil; record, recordErr = csvReader.Read() {
		var rev models.Revision
		rev.Member = new(models.Member)
		rev.Member.Affiliation = models.AffiliationNavy
		// The sailor's name, in the format LASTNAME,FIRSTNAME (optional MI) (optional suffix)
		fullname := record[fields["Service Member Name"]]
		names := strings.SplitN(fullname, ",", 2)
		rev.Member.FamilyName = names[0]
		names = strings.Fields(names[1])
		rev.Member.GivenName = names[0]
		if len(names) > 1 {
			if stringInSlice(names[len(names)-1], suffixes) {
				rev.Member.Suffix = &names[len(names)-1]
				if len(names) > 2 {
					middleName := strings.Join(names[1:len(names)-1], " ")
					rev.Member.MiddleName = &middleName
				}
			} else {
				middleName := strings.Join(names[1:], " ")
				rev.Member.MiddleName = &middleName
			}
		}

		daysStarting31Dec1899, _ := strconv.Atoi(record[fields["Order Create/Modification Date"]])
		dateIssued := time.Date(1899, time.December, 30+daysStarting31Dec1899, 0, 0, 0, 0, time.Local)
		fmtDateIssued := strfmt.DateTime(dateIssued)
		rev.DateIssued = &fmtDateIssued

		orderModNbr, orderModNbrErr := strconv.Atoi(record[fields["Order Modification Number"]])
		if orderModNbrErr != nil {
			orderModNbr = 0
		}
		obligModNbr, obligModNbrErr := strconv.Atoi(record[fields["Obligation Modification Number"]])
		if obligModNbrErr != nil {
			obligModNbr = 0
		}
		seqNum := int64(orderModNbr + obligModNbr)
		rev.SeqNum = &seqNum

		if record[fields["Obligation Status Code"]] == "D" {
			rev.Status = models.StatusCanceled
		} else {
			rev.Status = models.StatusAuthorized
		}
		rev.Member.Title = &record[fields["Rank Classification  Description"]]
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
			rev.LosingUnit.Name = &name
		}
		if uic := strings.TrimSpace(record[fields["Detach UIC"]]); len(uic) > 0 {
			fmtUIC := fmt.Sprintf("N%05s", uic)
			rev.LosingUnit.Uic = &fmtUIC
		}
		if city := strings.TrimSpace(record[fields["Detach UIC City Name"]]); len(city) > 0 {
			rev.LosingUnit.City = &city
		}
		if state := strings.TrimSpace(record[fields["Detach State Code"]]); len(state) > 0 {
			rev.LosingUnit.Locality = &state
		}
		if country := strings.TrimSpace(record[fields["Detach Country Code"]]); len(country) > 0 {
			rev.LosingUnit.Country = &country
		}

		daysStarting31Dec1899, daysError := strconv.Atoi(record[fields["Ultimate Estimated Arrival Date"]])
		if daysError == nil {
			estArrivalDate := time.Date(1899, time.December, 30+daysStarting31Dec1899, 0, 0, 0, 0, time.Local)
			rev.ReportNoLaterThan = new(strfmt.Date)
			*rev.ReportNoLaterThan = strfmt.Date(estArrivalDate)
		}

		rev.GainingUnit = new(models.Unit)
		if name := strings.TrimSpace(record[fields["Ultimate UIC Home Port"]]); len(name) > 0 {
			rev.GainingUnit.Name = &name
		}
		if uic := strings.TrimSpace(record[fields["Ultimate UIC"]]); len(uic) > 0 {
			fmtUIC := fmt.Sprintf("N%05s", uic)
			rev.GainingUnit.Uic = &fmtUIC
		}
		if city := strings.TrimSpace(record[fields["Ultimate UIC City Name"]]); len(city) > 0 {
			rev.GainingUnit.City = &city
		}
		if state := strings.TrimSpace(record[fields["Ultimate State Code"]]); len(state) > 0 {
			rev.GainingUnit.Locality = &state
		}
		if country := strings.TrimSpace(record[fields["Ultimate Country Code"]]); len(country) > 0 {
			rev.GainingUnit.Country = &country
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
		rev.PcsAccounting.Tac = &record[fields["TAC"]]

		var params operations.PostRevisionParams
		params.SetMemberID(record[fields["Ssn (obligation)"]])
		params.SetOrdersNum(record[fields["Primary SDN"]])
		params.SetIssuer(string(models.IssuerNavy))
		params.SetRevision(&rev)
		params.SetTimeout(time.Second * 30)
		_, err = ordersGateway.Operations.PostRevision(&params)
		if err != nil {
			log.Fatal(err)
		}

		bodyBuf := &bytes.Buffer{}
		encoder := json.NewEncoder(bodyBuf)
		encoder.SetIndent("", "  ")
		encoderErr := encoder.Encode(rev)
		if encoderErr != nil {
			log.Fatal(err)
		}

		fmt.Print(bodyBuf.String())
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
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
