package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// YamlConfig represents the koinos yaml application config values
type YamlConfig struct {
	Global     map[string]interface{} `yaml:"global,omitempty"`
	P2P        map[string]interface{} `yaml:"p2p,omitempty"`
	BlockStore map[string]interface{} `yaml:"block-store,omitempty"`
}

// GetStringOption fetches a string cli value, respecting values in a given config
func GetStringOption(key string, defaultValue string, cliArg string, configs ...map[string]interface{}) string {
	if cliArg != "" {
		return cliArg
	}

	for _, config := range configs {
		if v, ok := config[key]; ok {
			if option, ok := v.(string); ok {
				return option
			}
		}
	}

	return defaultValue
}

// GetStringSliceOption fetches a string slicecli value, respecting values in a given config
func GetStringSliceOption(key string, cliArg []string, configs ...map[string]interface{}) []string {
	stringSlice := cliArg

	for _, config := range configs {
		if v, ok := config[key]; ok {
			if slice, ok := v.([]interface{}); ok {
				for _, option := range slice {
					if str, ok := option.(string); ok {
						stringSlice = append(stringSlice, str)
					}
				}
			}
		}
	}

	return stringSlice
}

// InitYamlConfig initializes a yaml config
func InitYamlConfig(baseDir string) *YamlConfig {
	yamlConfigPath := filepath.Join(baseDir, "config.yml")
	if _, err := os.Stat(yamlConfigPath); os.IsNotExist(err) {
		yamlConfigPath = filepath.Join(baseDir, "config.yaml")
	}

	yamlConfig := YamlConfig{}
	if _, err := os.Stat(yamlConfigPath); err == nil {
		data, err := ioutil.ReadFile(yamlConfigPath)
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(data, &yamlConfig)
		if err != nil {
			panic(err)
		}
	} else {
		yamlConfig.Global = make(map[string]interface{})
		yamlConfig.P2P = make(map[string]interface{})
		yamlConfig.BlockStore = make(map[string]interface{})
	}

	return &yamlConfig
}
