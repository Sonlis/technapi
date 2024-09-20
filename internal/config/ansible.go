package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Hosts struct {
	Hosts map[string]interface{} `yaml:"hosts"`
}

func ParseAnsibleConfig(configPath string) (map[string]Hosts, error) {
	m := make(map[string]Hosts)
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal config file: %v", err)
	}

	return m, nil
}
