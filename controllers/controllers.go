package controllers

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func CheckDomain(domain string) {
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

	fmt.Printf("Domain= %v\n , hasMX= %v\n, hasSPF= %v\n, SPFRecord= %v\n, hasDMARC= %v\n, DMARCRecord= %v\n",
		domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
