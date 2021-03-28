package handlers

import (
	"testing"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
	"github.com/brandoncate-personal/monitor-cli/internal/core/services"
)

func Test_NewTestRunnerAdapter_Go_Success(t *testing.T) {

	config := domain.Config{
		Runner: string(services.Go),
	}

	testRunnerType, _, err := NewTestRunnerAdapter(config)
	if err != nil {
		t.Errorf("Test_NewTestRunnerAdapter_Go_Success failed, expected err to be nil, got %v", err)
	}
	if testRunnerType != services.Go {
		t.Errorf("Test_NewTestRunnerAdapter_Go_Success failed, expected testRunnerType to be Go, got %v", err)
	}
}
