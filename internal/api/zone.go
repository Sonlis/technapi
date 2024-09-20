package api

import (
	"fmt"
	"net/url"
	"os"
    "net/http"

	"gopkg.in/yaml.v3"
)

const (
    ZoneApiPath = "/api/zones"
)

type ZoneConfig struct {
    Zone string `yaml:"zone"`
    Type string `yaml:"type"`
    UseSoaSerialDateScheme string `yaml:"useSoaSerialDateScheme,omitempty"`
    PrimaryNameServerAddress string `yaml:"primaryNameServerAddresses,omitempty"`
    ZoneTransferProtocol string `yaml:"zoneTransferProtocol,omitempty"`
    TsigKeyName string `yaml:"tsigKeyName,omitempty"`
    Protocol string `yaml:"protocol,omitempty"`
    Forwarder string `yaml:"forwarder,omitempty"`
    DnssecValidation string `yaml:"dnssecValidation,omitempty"`
    ProxyType string `yaml:"proxyType,omitempty"`
    ProxyAddress string `yaml:"proxyAddress,omitempty"`
    ProxyPort string `yaml:"proxyPort,omitempty"`
    ProxyUsername string `yaml:"proxyUsername,omitempty"`
    ProxyPassword string `yaml:"proxyPassword,omitempty"`
}

type ZoneCreate struct {
    Response Zone `json:"response"`
}

type ZoneList struct {
    Response Zones `json:"response"`
}

 type Zones struct {
     PageNumber int `json:"pageNumber"`
     TotalPages int `json:"totalPages"`
     TotalZones int `json:"totalZones"`
     Zones []Zone `json:"zones"`
 }

 type Zone struct {
     Name string `json:"name,omitempty"`
     Type string `json:"type,omitempty"`
     IsExpired bool `json:"isExpired,omitempty"`
     Disabled bool `json:"disabled,omitempty"`
     Domain string `json:"domain,omitempty"`
 }

func ParseZoneConfig(configPath string) (*ZoneConfig, error) {
    yamlFile, err := os.ReadFile(configPath)
    if err != nil {
        return nil, fmt.Errorf("Failed to read config file: %v", err)
    }

    var zoneConfig *ZoneConfig
    err = yaml.Unmarshal(yamlFile, &zoneConfig)
    if err != nil {
        return nil, fmt.Errorf("Failed to unmarshal config file: %v", err)
    }

    return zoneConfig, nil
}

// This is so dumb why no JSON
// POST request is available but with form data, so in any case we would have to go through the same process
func (cfg *ZoneConfig) toQueryParameters() url.Values {
    v := url.Values{}
    v.Add("zone", cfg.Zone)
    v.Add("type", cfg.Type)
    if cfg.UseSoaSerialDateScheme != "" {
        v.Add("useSoaSerialDateScheme", cfg.UseSoaSerialDateScheme)
    }
    if cfg.PrimaryNameServerAddress != "" {
        v.Add("primaryNameServerAddress", cfg.PrimaryNameServerAddress)
    }
    if cfg.ZoneTransferProtocol != "" {
        v.Add("zoneTransferProtocol", cfg.ZoneTransferProtocol)
    }
    if cfg.TsigKeyName != "" {
        v.Add("tsigKeyName", cfg.TsigKeyName)
    }
    if cfg.Protocol != "" {
        v.Add("protocol", cfg.Protocol)
    }
    if cfg.Forwarder != "" {
        v.Add("forwarder", cfg.Forwarder)
    }
    if cfg.DnssecValidation != "" {
        v.Add("dnssecValidation", cfg.DnssecValidation)
    }
    if cfg.ProxyType != "" {
        v.Add("proxyType", cfg.ProxyType)
    }
    if cfg.ProxyAddress != "" {
        v.Add("proxyAddress", cfg.ProxyAddress)
    }
    if cfg.ProxyPort != "" {
        v.Add("proxyPort", cfg.ProxyPort)
    }
    if cfg.ProxyUsername != "" {
        v.Add("proxyUsername", cfg.ProxyUsername)
    }
    if cfg.ProxyPassword != "" {
        v.Add("proxyPassword", cfg.ProxyPassword)
    }
    return v
}

func (c *TechniClient) CreateZone(zoneConfig *ZoneConfig) (*ZoneCreate, error) {
    var z *ZoneCreate
    request_url := c.Url + ZoneApiPath + "/create"

    req, err := http.NewRequest("GET", request_url, nil)
    if err != nil {
        return nil, fmt.Errorf("Failed to initialize list zones request: %v", err)
    }

    queryParams := zoneConfig.toQueryParameters()
    req.URL.RawQuery = queryParams.Encode()

    c.setTokenQueryParam(req)

    respBody, err := c.executeRequest(req)
    if err != nil {
        return nil, fmt.Errorf("Failed to create Zone: %v", err)
    }

    err = yaml.Unmarshal(respBody, &z)
    if err != nil {
        return nil, fmt.Errorf("Failed to unmarshal Technitium's response: %v", err)
    }
    
    return z, nil
    
}

func (c *TechniClient) ListZones() (*ZoneList, error) {
    request_url := c.Url + ZoneApiPath + "/list"
    var z *ZoneList

    req, err := http.NewRequest("GET", request_url, nil)
    if err != nil {
        return nil, fmt.Errorf("Failed to initialize list zones request: %v", err)
    }

    c.setTokenQueryParam(req)

    respBody, err := c.executeRequest(req)
    if err != nil {
        return nil, fmt.Errorf("Failed to list Technitium zones: %v", err)
    }

    err = yaml.Unmarshal(respBody, &z)
    if err != nil {
        return nil, fmt.Errorf("Failed to unmarshal Technitium's response: %v", err)
    }
    
    return z, nil
}
