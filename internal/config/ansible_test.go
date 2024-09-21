package config

import (
	"os"
	"testing"
)

func TestParseAnsibleConfig(t *testing.T) {
	goodConf := []byte(`
rohan:
  hosts:
    101.100.159.12:
gondor:
  hosts:
    101.234.151.25:
    `)

	defer os.Remove("inventory.yaml")

	err := os.WriteFile("inventory.yaml", goodConf, 0644)
	if err != nil {
		t.Errorf("Failed to create dummy conf: %v", err)
	}

	hosts, err := ParseAnsibleConfig("./inventory.yaml")
	if err != nil {
		t.Error(err)
	}

	if _, ok := hosts["rohan"]; !ok {
		t.Errorf("Rohan was not found in the hosts list.")
	}

	if _, ok := hosts["gondor"]; !ok {
		t.Errorf("Where was gondor when the whestfold fell?")
	}

	if _, ok := hosts["rohan"].Hosts["101.100.159.12"]; !ok {
		t.Errorf("Rohan's address was not found")
	}

	if _, ok := hosts["gondor"].Hosts["101.234.151.25"]; !ok {
		t.Errorf("Gondor's address was not found")
	}

	badConf := []byte(`
mordor:
  eye:
    ever-watchful:
isengard:
  hosts:
    they-taking-the-hobitts-to-isengard:
    `)
	err = os.WriteFile("inventory.yaml", badConf, 0644)
	if err != nil {
		t.Errorf("Failed to create dummy conf: %v", err)
	}

	hosts, err = ParseAnsibleConfig("./inventory.yaml")
	if err != nil {
		t.Error(err)
	}

	if _, ok := hosts["mordor"]; !ok {
		t.Errorf("Mordor was not found in the hosts list.")
	}

	if _, ok := hosts["isengard"]; !ok {
		t.Errorf("Isengard was not found in the hosts list.")
	}

	if _, ok := hosts["mordor"].Hosts["whatever"]; ok {
		t.Errorf("Mordor should not have any hosts")
	}

	// There is no checks whether the input is a valid IP address.
	// So the isengard conf is fine as far as the parseAnsibleConfig function goes.
}
