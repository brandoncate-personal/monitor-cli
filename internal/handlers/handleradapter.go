package handlers

import (
	"errors"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
	"github.com/brandoncate-personal/monitor-cli/internal/core/services"
)

func NewTestRunnerAdapter(config domain.Config) (services.TestRunnerType, interface{}, error) {
	if config.Runner == string(services.Go) {
		return "go", services.GoTestRunner{}, nil
	}

	return "", nil, errors.New("unsupported test runner type passed")
}
