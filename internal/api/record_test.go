package api

import (
	"testing"

	"github.com/sonlis/technapi/internal/config"
)

func TestRemoveRootDomain(t *testing.T) {
	// There is nowhere in the code where it is checked if the record is a valid domain name,
	// neither if it's actually a root domain. In the current use case there's no need to,
	// but could need to check in the future.
	records := []Record{
		{
			Name: "botlane.rift.riot",
			Type: "A",
			RData: RecordData{
				IpAddress: "1.1.1.1",
			},
		},
		{
			Name: "midlane.rift.riot",
			Type: "A",
			RData: RecordData{
				IpAddress: "2.2.2.2",
			},
		},
	}
	for i := range records {
		records[i].removeRootDomain()
	}

	if records[0].Name != "botlane" {
		t.Errorf("Error remove root domain: want %s, got %s", "botlane", records[0].Name)
	}

	if records[1].Name != "midlane" {
		t.Errorf("Error remove root domain: want %s, got %s", "midlane", records[1].Name)
	}
}

func TestCreateRecord(t *testing.T) {
	// Weird to create a zone in this function to test if records exist
	// It should be done by the test executor prior to the tests
	c, err := GetTestClient()
	if err != nil {
		t.Error(err)
	}

	record := config.DnsRecord{
		Record: "botlane",
		Ip:     "3.3.3.3",
	}

	zone := &config.ZoneConfig{
		Zone: "rift.riot",
		Type: "Primary",
	}

	_, err = c.CreateZone(zone)
	if err != nil {
		t.Errorf("Error creating zone: %v", err)
	}

	err = c.CreateRecord(record, zone.Zone)
	if err != nil {
		t.Errorf("Error creating record: %v", err)
	}

	record = config.DnsRecord{
		Record: "botlane",
		Ip:     "ashe",
	}

	err = c.CreateRecord(record, zone.Zone)
	if err == nil {
		t.Errorf("Should get an error of creating a wrong record, but got none")
	}
}

func TestGetRecords(t *testing.T) {
	// Weird to create a zone in this function to test if records exist
	// It should be done by the test executor prior to the tests
	c, err := GetTestClient()
	if err != nil {
		t.Error(err)
	}

	record := config.DnsRecord{
		Record: "jungle",
		Ip:     "4.4.4.4",
	}

	zone := &config.ZoneConfig{
		Zone: "rift.lol",
		Type: "Primary",
	}

	_, err = c.CreateZone(zone)
	if err != nil {
		t.Errorf("Error creating zone: %v", err)
	}

	// Same, the records should be created prior execution
	err = c.CreateRecord(record, zone.Zone)
	if err != nil {
		t.Errorf("Error creating record: %v", err)
	}

	records, err := c.GetRecords("rift.lol")
	if err != nil {
		t.Errorf("Error getting records: %v", err)
	}

	rExist := false
	for _, r := range records {
		if r.RData.IpAddress == record.Ip && r.Name == record.Record {
			rExist = true
		}
	}

	if !rExist {
		t.Errorf("The created records could not be found.")
	}
}
