package controllers

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

	"github.com/sahandPgr/Email-verifier/models"
)

// Checking domain is valid or not
func CheckDomain(mail string) models.Response {
	res, _ := regexp.Compile(`([a-zA-Z0-9.-]+)$`)
	domain := res.FindString(mail)
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	fmt.Println("Checking domain for: " + domain)
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	domainResponse := models.Response{
		Domain:      domain,
		HasMX:       hasMx,
		HasSPF:      hasSPF,
		SPFRecord:   spfRecord,
		HasDMARC:    hasDMARC,
		DMARCRecord: dmarcRecord,
	}

	fmt.Printf("Domain= %v\n , hasMX= %v\n, hasSPF= %v\n, SPFRecord= %v\n, hasDMARC= %v\n, DMARCRecord= %v\n",
		domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)

	return domainResponse
}
