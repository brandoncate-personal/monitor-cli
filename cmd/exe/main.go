package main

import (
	"log"

	"github.com/brandoncate-personal/monitor-cli/internal/core/services"
	"github.com/brandoncate-personal/monitor-cli/internal/handlers"
	"github.com/brandoncate-personal/monitor-cli/internal/repositories"
)

func main() {
	// parse argument flags -- specific to cmd/exe package
	args := ParseArgs()

	// load meta monitor info based on flags passed -- handler
	handlerType, handlerArgs, err := NewHandlerAdapter(args)
	if err != nil {
		log.Fatalf("Error occurred in NewHandlerTranslator: %v", err)
	}
	configLoader, err := handlers.NewHandler(handlerType, handlerArgs)
	if err != nil {
		log.Fatalf("Error occurred in NewHandler: %v", err)
	}
	config, err := configLoader.Load()
	if err != nil {
		log.Fatalf("Error occurred in Config Loader Load: %v", err)
	}

	// Load go test runner service -- go core service
	testRunnerType, testRunnerArgs, err := handlers.NewTestRunnerAdapter(config)
	if err != nil {
		log.Fatalf("Error occurred in NewTestRunnerTranslator: %v", err)
	}
	testRunner, err := services.NewTestRunner(testRunnerType, testRunnerArgs)
	if err != nil {
		log.Fatalf("Error occurred in NewTestRunner: %v", err)
	}

	// Load Http Exporter -- repository
	exporterType, exporterArgs, err := NewRepositoryAdapter(config)
	if err != nil {
		log.Fatalf("Error occurred in NewRepositoryAdapter: %v", err)
	}
	exporter, err := repositories.NewExporter(exporterType, exporterArgs)
	if err != nil {
		log.Fatalf("Error occurred in NewExporter: %v", err)
	}

	// inject dynamically loaded test runner and exporter to executor service
	executor := services.NewTestExecutor(testRunner, exporter)

	// load scheduler
	testScheduler := services.NewTestScheduler()

	for _, test := range config.Tests {
		testScheduler.Schedule(test.Schedule, *executor, test.Test)
	}

	// execute scheduler
	testScheduler.Start()
}
