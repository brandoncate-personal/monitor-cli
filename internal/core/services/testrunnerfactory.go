package services

import (
	"errors"

	"github.com/brandoncate-personal/monitor-cli/internal/core/ports"
)

type TestRunnerType string

const (
	Go TestRunnerType = "go"
)

// Here, NewPerson returns an interface, and not the person struct itself
func NewTestRunner(testRunnerType TestRunnerType, params interface{}) (ports.TestRunnerService, error) {

	switch testRunnerType {
	case Go:
		// no required params
		return GoTestRunner{}, nil
	default:
		return nil, errors.New("unsupported test runnerf   nz type")
	}

}
