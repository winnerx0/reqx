package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

func Parse(path string) (*Config, error) {

	ext := filepath.Ext(path)

	if ext != ".yaml" && ext != ".yml" {
		return &Config{}, errors.New("Invalid file format. Only YAML allowed.")
	}

	var config Config

	fileBytes, err := os.ReadFile(path)

	if err != nil {
		return &Config{}, errors.New("reqx.yaml not found in current directory")
	}

	err = yaml.Unmarshal(fileBytes, &config)

	if err != nil {
		return &Config{}, err
	}

	return &config, nil
}
