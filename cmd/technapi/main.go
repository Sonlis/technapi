package main

import (
	"fmt"
	"os"

	"github.com/sonlis/technapi/internal/api"
	"github.com/sonlis/technapi/internal/config"
)

func main() {
	technitiumUrl, user, password, err := setupVariables()
	if err != nil {
		fmt.Printf("Missing environment variable: %v\n", err)
		os.Exit(1)
	}

	client := api.TechniClient{
		Url: technitiumUrl,
	}

	err = client.GetSessionToken(user, password)
	if err != nil {
		fmt.Printf("Error retrieving session token: %v\n", err)
		os.Exit(1)
	}

	zones, err := client.ListZones()
	if err != nil {
		fmt.Printf("Error listing zones: %v\n", err)
		os.Exit(1)
	}
    
	zoneConfig, err := config.ParseZoneConfig("./zone-config.yaml")
	if err != nil {
		fmt.Printf("Error parsing local zone configuration: %v\n", err)
		os.Exit(1)
	}

	createZone := true
	for _, zone := range zones {
		if zone.Name == zoneConfig.Zone {
			if zone.Type == zoneConfig.Type {
				createZone = false
			}
		}
	}

	if createZone {
		_, err := client.CreateZone(zoneConfig)
		if err != nil {
			fmt.Printf("Error creating zone: %v\n", err)
			os.Exit(1)
		}
	}

	records, err := client.GetRecords(zoneConfig.Zone)
	if err != nil {
		fmt.Printf("Error getting records: %v\n", err)
		os.Exit(1)
	}

	ansibleConfig, err := config.ParseAnsibleConfig("./inventory.yaml")
	if err != nil {
		fmt.Printf("Error parsing ansible configuration: %v\n", err)
		os.Exit(1)
	}

	technitiumRecords := parseZoneRecords(records)
	ansibleRecords := parseAnsibleHosts(ansibleConfig)
	recordsToCreate := filterExistingRecords(technitiumRecords, ansibleRecords)

	for _, record := range recordsToCreate {
		err := client.CreateRecord(record, zoneConfig.Zone)
		if err != nil {
			fmt.Printf("Error creating record %s: %v\n", record.Record, err)
		}
	}
}

func parseZoneRecords(records []api.Record) []config.DnsRecord {
	rs := []config.DnsRecord{}
	for _, record := range records {
		if record.RData.IpAddress != "" {
			r := config.DnsRecord{
				Ip:     record.RData.IpAddress,
				Record: record.Name,
			}
			rs = append(rs, r)
		}
	}
	return rs
}

func parseAnsibleHosts(c map[string]config.Hosts) []config.DnsRecord {
	rs := []config.DnsRecord{}
	for record, data := range c {
		for ip := range data.Hosts {
			r := config.DnsRecord{
				Ip:     ip,
				Record: record,
			}
			rs = append(rs, r)
		}
	}
	return rs
}

func filterExistingRecords(techRecord, ansibleRecord []config.DnsRecord) []config.DnsRecord {
	recordsToCreate := []config.DnsRecord{}
	for _, ar := range ansibleRecord {
		exist := false
		for _, tr := range techRecord {
			if ar.Record == tr.Record && ar.Ip == tr.Ip {
				exist = true
				break
			}
		}
		if !exist {
			recordsToCreate = append(recordsToCreate, ar)
		}
	}
	return recordsToCreate
}

func setupVariables() (string, string, string, error) {
	url := os.Getenv("TECHNITIUM_URL")
	if url == "" {
		return "", "", "", fmt.Errorf("TECHNITIUM_URL not set.")
	}
	user := os.Getenv("TECHNITIUM_USER")
	if user == "" {
		return "", "", "", fmt.Errorf("TECHNITIUM_USER not set.")
	}
	password := os.Getenv("TECHNITIUM_PASSWORD")
	if password == "" {
		return "", "", "", fmt.Errorf("TECHNITIUM_PASSWORD not set.")
	}
	return url, user, password, nil
}
