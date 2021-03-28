package main

import (
	"errors"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
	"github.com/brandoncate-personal/monitor-cli/internal/repositories"
)

func NewRepositoryAdapter(config domain.Config) (repositories.ExporterType, interface{}, error) {
	if config.Export.Http != "" {
		return "http", repositories.HttpExporter{
			Url: config.Export.Http,
		}, nil
	}

	return "", nil, errors.New("unsupported repository passed")
}
