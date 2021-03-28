package ports

import "github.com/brandoncate-personal/monitor-cli/internal/core/domain"

type ExporterService interface {
	Export(domain.TestResults) error
}
