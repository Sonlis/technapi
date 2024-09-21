package api

import (
	"net/http"
	"net/url"
	"testing"
)


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
