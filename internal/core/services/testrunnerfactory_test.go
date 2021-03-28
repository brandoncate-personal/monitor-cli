package services

import (
	"testing"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
)

func Test_TestRunnerFactory_Go_Success(t *testing.T) {

	test := domain.Test{
		Dir:  "faketests",
		File: "go_fake_test.go",
		Name: "TestCase1",
	}

	goTestRunner, err := NewTestRunner(Go, nil)
	if err != nil {
		t.Errorf("Test_GoTestRunner_Success failed, expected err to be nil, got %v", err)
	}

	testResults, err := goTestRunner.Run(test)
	if err != nil {
		t.Errorf("Test_GoTestRunner_Success failed, expected err to be nil, got %v", err)
	}
	if testResults.StatusCode != domain.OK {
		t.Errorf("Test_GoTestRunner_Success failed, expected StatusCode to be 200, got %v", testResults.StatusCode)
	}
}
