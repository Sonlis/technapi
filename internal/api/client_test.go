package api

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

const (
	testUser     = "admin"
	testPassword = "admin"
	testUrl      = "http://localhost:5380"
)

func GetTestClient() (*TechniClient, error) {
	c := &TechniClient{
		Url: testUrl,
	}

	err := c.GetSessionToken(testUser, testPassword)
	if err != nil {
		return nil, fmt.Errorf("Error getting session token from Technitium: %v", err)
	}

	if c.sessionToken == "" {
		return nil, fmt.Errorf("Session token is empty")
	}

	return c, nil
}

func TestGetSessionToken(t *testing.T) {
	c := TechniClient{
		Url: testUrl,
	}
	err := c.GetSessionToken(testUser, testPassword)
	if err != nil {
		t.Errorf("Error getting session token: %v", err)
	}
	if c.sessionToken == "" {
		t.Errorf("Failed to get session token from technitium")
	}
}

func TestSetTokenQueryParam(t *testing.T) {
	c, err := GetTestClient()
	if err != nil {
		t.Error(err)
	}

	request_url := c.Url + ApiUserPath + "/login"

	want := url.Values{
		"token": []string{c.sessionToken},
	}

	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		t.Errorf("Failed to initialize get token request: %v", err)
	}

	got := c.setTokenQueryParam(req)
	if want["token"][0] != got["token"][0] {
		t.Errorf("Setting the token in query parameters failed: want %v, got %v", want, got)
	}
}
