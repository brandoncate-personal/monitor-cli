package services

import (
	"testing"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
)

func Test_GoTestRunner_Pass_Success(t *testing.T) {

	test := domain.Test{
		Dir:  "data",
		File: "fake_test.go",
		Name: "Test_Pass",
	}

	goTestRunner := GoTestRunner{}
	testResults, err := goTestRunner.Run(test)

	if err != nil {
		t.Errorf("Test_GoTestRunner_Success failed, expected err to be nil, got %v", err)
	}
	if testResults.StatusCode != domain.OK {
		t.Errorf("Test_GoTestRunner_Success failed, expected StatusCode to be 200, got %v", testResults.StatusCode)
	}
}

func Test_GoTestRunner_Random_No_Cache_Success(t *testing.T) {

	test := domain.Test{
		Dir:  "data",
		File: "fake_test.go",
		Name: "Test_Random",
	}

	goTestRunner := GoTestRunner{}
	_, err := goTestRunner.Run(test)

	if err != nil {
		t.Errorf("Test_GoTestRunner_Success failed, expected err to be nil, got %v", err)
	}
}
