package main

import (
	"fmt"

	"github.com/sonlis/technapi/internal/api"
	"github.com/sonlis/technapi/internal/config"
)


func main() {
	client := api.TechniClient{
		Url: "http://192.168.0.13:5380",
	}
	client.GetSessionToken("", "")
	zones, err := client.ListZones()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	zoneConfig, err := config.ParseZoneConfig("./zone-config.yaml")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	createZone := true
	for _, zone := range zones.Response.Zones {
		if zone.Name == zoneConfig.Zone {
			if zone.Type == zoneConfig.Type {
				createZone = false
			}
		}
	}

	if createZone {
		_, err := client.CreateZone(zoneConfig)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	records, err := client.GetRecords(zoneConfig.Zone)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	ansibleConfig, err := config.ParseAnsibleConfig("./inventory.yaml")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	technitiumRecords := parseZoneRecords(records)
	ansibleRecords := parseAnsibleHosts(ansibleConfig)
	recordsToCreate := filterExistingRecords(technitiumRecords, ansibleRecords)

    for _, record := range recordsToCreate {
        err := client.CreateRecord(record, zoneConfig.Zone)
        if err != nil {
            fmt.Println("Error: ", err)
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
