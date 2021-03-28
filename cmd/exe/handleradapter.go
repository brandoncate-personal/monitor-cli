package main

import (
	"errors"

	"github.com/brandoncate-personal/monitor-cli/internal/handlers"
)

func NewHandlerAdapter(args Arguments) (handlers.HandlerType, interface{}, error) {
	if args.yaml != "" {
		return "yaml", handlers.YamlConfigLoader{
			Path: args.yaml,
		}, nil
	}

	return "", nil, errors.New("unsupported handler passed")
}
