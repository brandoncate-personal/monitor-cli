package handlers

import (
	"io/ioutil"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
	"gopkg.in/yaml.v2"
)

type YamlConfigLoader struct {
	Path string
}

func (loader YamlConfigLoader) Load() (domain.Config, error) {
	config := domain.Config{}

	yamlFile, err := ioutil.ReadFile(loader.Path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
