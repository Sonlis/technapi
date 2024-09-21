package config

import (
	"os"
	"testing"
)

func TestParseZoneConfig(t *testing.T) {
	goodConf := []byte(`
zone: hogwarts.scotland
type: Primary
    `)

	defer os.Remove("zone-config.yaml")

	err := os.WriteFile("zone-config.yaml", goodConf, 0644)
	if err != nil {
		t.Errorf("Failed to create dummy conf: %v", err)
	}

	zoneConfig, err := ParseZoneConfig("zone-config.yaml")
	if err != nil {
		t.Errorf("Failed to open config file: %v", err)
	}

	if zoneConfig.Zone != "hogwarts.scotland" {
		t.Errorf("Got wrong zone, expected hogwarts.scotland, got %v", err)
	}
	if zoneConfig.Type != "Primary" {
		t.Errorf("Got wrong zone type, expected Primary, got %v", err)
	}

	badConf := []byte(`
zonecaditquoi: hogwarts.scotland
firetype: Primary
    `)

	err = os.WriteFile("zone-config.yaml", badConf, 0644)
	if err != nil {
		t.Errorf("Failed to create dummy conf: %v", err)
	}

	// Useless to test here, there is no check (yet?) to see if the conf is proper.
}

func TestToQueryParameters(t *testing.T) {
	zc := *&ZoneConfig{
		Zone:          "zone",
		Type:          "Primary",
		ProxyUsername: "oui",
		ProxyPassword: "non",
	}

	got := zc.ToQueryParameters()
	want := map[string][]string{
		"zone":          []string{"zone"},
		"type":          []string{"Primary"},
		"proxyUsername": []string{"oui"},
		"proxyPassword": []string{"non"},
	}
	if want["zone"][0] != got["zone"][0] {
		t.Errorf("Error getting parameters from config: got %v, want %v", got["zone"][0], want["zone"][0])
	}
	if want["type"][0] != got["type"][0] {
		t.Errorf("Error getting parameters from config: got %v, want %v", got["type"][0], want["type"][0])
	}
	if want["proxyUsername"][0] != got["proxyUsername"][0] {
		t.Errorf("Error getting parameters from config: got %v, want %v", got["proxyUsername"][0], want["proxyUsername"][0])
	}
	if want["proxyPassword"][0] != got["proxyPassword"][0] {
		t.Errorf("Error getting parameters from config: got %v, want %v", got["proxyPassword"][0], want["proxyPassword"][0])
	}
}
