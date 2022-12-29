package domain_checker

import (
	"log"
	"net"
	"strings"
)

type Domain struct {
	Name        string
	HasMX       bool
	HasSPF      bool
	HasDMARC    bool
	SpfRecord   string
	DmarcRecord string
}

func checkDomain(domainName string) *Domain {
	domain := &Domain{}

	mxRecords, err := net.LookupMX(domainName)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	if len(mxRecords) > 0 {
		domain.HasMX = true
	}

	txtRecords, err := net.LookupTXT(domainName)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			domain.HasSPF = true
			domain.SpfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domainName)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			domain.HasDMARC = true
			domain.DmarcRecord = record
			break
		}
	}

	return domain
}
