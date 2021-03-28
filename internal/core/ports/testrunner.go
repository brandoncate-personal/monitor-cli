package ports

import "github.com/brandoncate-personal/monitor-cli/internal/core/domain"

type TestRunnerService interface {
	Run(test domain.Test) (domain.TestResults, error)
}
