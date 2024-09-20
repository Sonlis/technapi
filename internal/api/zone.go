package api

import (
	"encoding/json"
	"fmt"
	"github.com/sonlis/technapi/internal/config"
	"net/http"
)

const (
	ZoneApiPath = "/api/zones"
)

type ZoneCreate struct {
	Response Zone `json:"response"`
}

type ZoneList struct {
	Response Zones `json:"response"`
}

type Zones struct {
	PageNumber int    `json:"pageNumber"`
	TotalPages int    `json:"totalPages"`
	TotalZones int    `json:"totalZones"`
	Zones      []Zone `json:"zones"`
}

type Zone struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	IsExpired bool   `json:"isExpired,omitempty"`
	Disabled  bool   `json:"disabled,omitempty"`
	Domain    string `json:"domain,omitempty"`
}

func (c *TechniClient) CreateZone(zoneConfig *config.ZoneConfig) (*ZoneCreate, error) {
	var z *ZoneCreate
	request_url := c.Url + ZoneApiPath + "/create"

	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize list zones request: %v", err)
	}

	queryParams := zoneConfig.ToQueryParameters()
	req.URL.RawQuery = queryParams.Encode()

	c.setTokenQueryParam(req)

	respBody, err := c.executeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Zone: %v", err)
	}

	err = json.Unmarshal(respBody, &z)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal Technitium's response: %v", err)
	}

	return z, nil
}

func (c *TechniClient) ListZones() (*ZoneList, error) {
	var z *ZoneList
	request_url := c.Url + ZoneApiPath + "/list"

	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize list zones request: %v", err)
	}

	c.setTokenQueryParam(req)

	respBody, err := c.executeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to list Technitium zones: %v", err)
	}

	err = json.Unmarshal(respBody, &z)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal Technitium's response: %v", err)
	}

	return z, nil
}
