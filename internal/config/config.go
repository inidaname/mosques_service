package config

import (
	"log"
	"os"

	"github.com/inidaname/mosque/mosques-service/internal/types"
	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*types.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open config file: %v", err)
		return nil, err
	}
	defer file.Close()

	cfg := &types.Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		log.Fatalf("could not decode config YAML: %v", err)
		return nil, err
	}

	return cfg, nil
}
