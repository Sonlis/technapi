package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sonlis/technapi/internal/config"
)

const (
	recordApiPath = "/api/zones/records"
)

type Record struct {
	Name  string     `json:"name"`
	Type  string     `json:"type"`
	Ttl   int        `json:"ttl"`
	RData RecordData `json:"rData"`
}

type RecordData struct {
	NameServer string `json:"nameserver"`
	IpAddress  string `json:"ipAddress"`
}

type RecordList struct {
	Response Records `json:"response"`
}

type Records struct {
	Records []Record `json:"records"`
}

func (c *TechniClient) GetRecords(domain string) ([]Record, error) {
	var r *RecordList
	request_url := c.Url + recordApiPath + "/get"

	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize list zones request: %v", err)
	}

	q := setGetRecordsParams(domain, req)
	req.URL.RawQuery = q.Encode()

	c.setTokenQueryParam(req)

	respBody, err := c.executeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Zone: %v", err)
	}

	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal Technitium's response: %v", err)
	}

	records := removeRootDomain(r.Response.Records)

	return records, nil
}

func (c *TechniClient) CreateRecord(r config.DnsRecord, zone string) error {
	request_url := c.Url + recordApiPath + "/add"

	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		return fmt.Errorf("Failed to initialize add record request: %v", err)
	}

	q := setCreateRecordParams(r, zone, req)
	req.URL.RawQuery = q.Encode()

	c.setTokenQueryParam(req)

	respBody, err := c.executeRequest(req)
	if err != nil {
		return fmt.Errorf("Failed to add record: %v", err)
	}
    fmt.Println(req.URL)
    fmt.Println(string(respBody))

   return nil 
}

func removeRootDomain(records []Record) []Record {
	for i := range records {
		records[i].Name = strings.Split(records[i].Name, ".")[0]
	}
	return records
}

func setGetRecordsParams(domain string, req *http.Request) url.Values {
	q := req.URL.Query()
	q.Add("domain", domain)
	q.Add("listZone", "true")
	return q
}

func setCreateRecordParams(record config.DnsRecord, zone string, req *http.Request) url.Values {
    q := req.URL.Query()
    domain := record.Record + "." + zone
    q.Add("zone", zone)
    q.Add("domain", domain)
    q.Add("type", "A")
    q.Add("ipAddress", record.Ip)
    return q
}

