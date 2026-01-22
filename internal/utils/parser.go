package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
)

type Config struct {
	Requests []Request `yaml:"requests"`
}

type Request struct {
	Name    string            `yaml:"name"`
	Url     string            `yaml:"url"`
	Body    any               `yaml:"body"`
	Headers map[string]string `yaml:"headers"`
	Method  Method            `yaml:"method"`
}

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
