package services

import (
	"log"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
	"github.com/brandoncate-personal/monitor-cli/internal/core/ports"
)

type TestExecutor struct {
	testRunner ports.TestRunnerService
	exporter   ports.ExporterService
}

func NewTestExecutor(testRunner ports.TestRunnerService, exporter ports.ExporterService) *TestExecutor {
	return &TestExecutor{
		testRunner: testRunner,
		exporter:   exporter,
	}
}

func (executor *TestExecutor) Execute(test domain.Test) {
	testResults, err := executor.testRunner.Run(test)
	if err != nil { // do not exit on failure for scheduler
		log.Printf("error occurred in test runner run method: %v", err)
	}

	err = executor.exporter.Export(testResults)
	if err != nil { // do not exit on failure for scheduler
		log.Printf("error occurred in exporter export method: %v", err)
	}

}
