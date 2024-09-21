package api

import (
	"testing"

	"github.com/sonlis/technapi/internal/config"
)

func TestCreateZone(t *testing.T) {
	c, err := GetTestClient()
	if err != nil {
		t.Error(err)
	}
	z := &config.ZoneConfig{
		Zone: "testing-zone",
		Type: "Primary",
	}

	zCreate, err := c.CreateZone(z)
	if err != nil {
		t.Errorf("Error creating zone: %v", err)
	}

	if zCreate.Domain != z.Zone {
		t.Errorf("Got the wrong domain name in response: want %s, got %s", z.Zone, zCreate.Domain)
	}

	z.Zone = "not-good-type"
	z.Type = "FIRE"
	zCreate, err = c.CreateZone(z)
	if err == nil {
		t.Error("Should get an error, but got none")
	}
}

func TestListZones(t *testing.T) {
	c, err := GetTestClient()
	if err != nil {
		t.Error(err)
	}

	zone1 := &config.ZoneConfig{
		Zone: "zone1",
		Type: "Primary",
	}
	_, err = c.CreateZone(zone1)
	if err != nil {
		t.Errorf("Failed to create test zone 1: %v", err)
	}

	zone2 := &config.ZoneConfig{
		Zone: "zone2",
		Type: "Primary",
	}
	_, err = c.CreateZone(zone2)
	if err != nil {
		t.Errorf("Failed to create test zone 2: %v", err)
	}

	zones, err := c.ListZones()
	if err != nil {
		t.Errorf("Failed to list zones: %v", err)
	}

	z1Exist, z2Exist := false, false
	for _, zone := range zones {
		if zone1.Zone == zone.Name {
			z1Exist = true
		} else if zone2.Zone == zone.Name {
			z2Exist = true
		}
	}

	if !z1Exist || !z2Exist {
		t.Errorf("Zones created were not found in listing zones response: %v", zones)
	}
}
