package configuration

import (
	"gopkg.in/yaml.v2"
	"os"
)

const (
	ConfigPath = "./resources/mongo-config.yml"
)

type MongoConfig struct {
	URL        string `yaml:"url"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

func ReadMongoConfig() (*MongoConfig, error) {
	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}

	var config MongoConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
