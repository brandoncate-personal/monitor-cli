package handlers

import (
	"errors"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
)

type ConfigLoader interface {
	Load() (domain.Config, error)
}

type HandlerType string

const (
	Yaml HandlerType = "yaml"
)

// Here, NewPerson returns an interface, and not the person struct itself
func NewHandler(handlerType HandlerType, params interface{}) (ConfigLoader, error) {

	switch handlerType {
	case Yaml:
		yamlParams, ok := params.(YamlConfigLoader)
		if !ok {
			return nil, errors.New("required params for yaml loader not passed")
		}
		return YamlConfigLoader{
			Path: yamlParams.Path,
		}, nil
	default:
		return nil, errors.New("unsupported handler type")
	}

}
