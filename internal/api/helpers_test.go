package api

import (
	"net/http"
	"testing"
)

func TestCheckResponseStatus(t *testing.T) {
	body := []byte(`{"status": "ok", "errorMessage": ""}`)
	err := checkResponseStatus(body)
	if err != nil {
		t.Errorf("Should not get an error but got %v", err)
	}

	body = []byte(`{"status": "error", "errorMessage": "You did something very wrong"}`)
	err = checkResponseStatus(body)
	if err == nil {
		t.Errorf("Should get an error but got none")
	}
}

func TestExecuteRequest(t *testing.T) {
	c, err := GetTestClient()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", testUrl+"/api/zones/list", nil)
	if err != nil {
		t.Errorf("Error building test execute request: %v", err)
	}
	c.setTokenQueryParam(req)

	_, err = c.executeRequest(req)
	if err != nil {
		t.Errorf("Error executing request: %v", err)
	}

	// Querying a wrong path from technitium should not return an error
	req, err = http.NewRequest("GET", testUrl+"/not/a/technitium/path", nil)
	if err != nil {
		t.Errorf("Error building test execute request: %v", err)
	}
	c.setTokenQueryParam(req)

	_, err = c.executeRequest(req)
	if err != nil {
		t.Errorf("Error executing request: %v", err)
	}

	// Requesting a wrongly formatted host should return an error
	req, err = http.NewRequest("GET", "this-is-not-a-host"+"/not/a/technitium/path", nil)
	if err != nil {
		t.Errorf("Error building test execute request: %v", err)
	}
	c.setTokenQueryParam(req)

	_, err = c.executeRequest(req)
	if err == nil {
		t.Error("Requesting a wrong host should return an error, but it didn't")
	}
}
