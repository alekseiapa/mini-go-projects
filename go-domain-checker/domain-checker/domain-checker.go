package domain_checker

import (
	"net"
	"regexp"
	"strings"
)

type Domain struct {
	Name        string
	IsValid     bool
	HasMX       bool
	HasSPF      bool
	HasDMARC    bool
	MXRecord    string
	SpfRecord   string
	DmarcRecord string
}

func isValidDomain(domain string) bool {
	// Define a regular expression that matches domain names
	domainRegexp := regexp.MustCompile(`^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$`)

	// Check if the domain name matches the regular expression
	return domainRegexp.MatchString(domain)
}

func CheckDomain(domainName string) *Domain {
	domain := &Domain{
		Name: domainName,
	}
	if !isValidDomain(domainName) {
		domain.IsValid = false
		return domain
	}
	domain.IsValid = true
	mxRecords, err := net.LookupMX(domainName)
	if err != nil {
		return domain
	}
	if len(mxRecords) > 0 {
		domain.MXRecord = mxRecords[0].Host
		domain.HasMX = true
	}

	txtRecords, err := net.LookupTXT(domainName)
	if err != nil {
		return domain
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
		return domain
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
