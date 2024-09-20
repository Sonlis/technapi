package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
)

type ZoneConfig struct {
	Zone                     string `yaml:"zone"`
	Type                     string `yaml:"type"`
	UseSoaSerialDateScheme   string `yaml:"useSoaSerialDateScheme,omitempty"`
	PrimaryNameServerAddress string `yaml:"primaryNameServerAddresses,omitempty"`
	ZoneTransferProtocol     string `yaml:"zoneTransferProtocol,omitempty"`
	TsigKeyName              string `yaml:"tsigKeyName,omitempty"`
	Protocol                 string `yaml:"protocol,omitempty"`
	Forwarder                string `yaml:"forwarder,omitempty"`
	DnssecValidation         string `yaml:"dnssecValidation,omitempty"`
	ProxyType                string `yaml:"proxyType,omitempty"`
	ProxyAddress             string `yaml:"proxyAddress,omitempty"`
	ProxyPort                string `yaml:"proxyPort,omitempty"`
	ProxyUsername            string `yaml:"proxyUsername,omitempty"`
	ProxyPassword            string `yaml:"proxyPassword,omitempty"`
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
func (cfg *ZoneConfig) ToQueryParameters() url.Values {
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
