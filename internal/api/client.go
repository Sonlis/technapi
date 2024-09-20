package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type TechniClient struct {
	Url          string
	sessionToken string
	c            http.Client
}

type LoginResponse struct {
	Token string `json:"token"`
}

const (
	ApiUserPath = "/api/user"
)

func (c *TechniClient) GetSessionToken(username, password string) error {
	request_url := c.Url + ApiUserPath + "/login"

	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		return fmt.Errorf("Failed to initialize get token request: %v", err)
	}

	q := req.URL.Query()
	q.Add("user", username)
	q.Add("pass", password)
	req.URL.RawQuery = q.Encode()

	var loginResponse *LoginResponse

	respBody, err := c.executeRequest(req)

	err = json.Unmarshal(respBody, &loginResponse)
	if err != nil {
		return fmt.Errorf("Failed to get token from Technitium response: %v", err)
	}

	c.sessionToken = loginResponse.Token
	return nil
}

func (c *TechniClient) setTokenQueryParam(req *http.Request) url.Values {
	q := req.URL.Query()
	q.Add("token", c.sessionToken)
	req.URL.RawQuery = q.Encode()
	return q
}
