package services

import (
	"errors"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
	"github.com/jasonlvhit/gocron"
)

type TestScheduler struct {
	scheduler *gocron.Scheduler
}

func NewTestScheduler() *TestScheduler {
	return &TestScheduler{
		scheduler: gocron.NewScheduler(),
	}
}

func (testScheduler *TestScheduler) Schedule(schedule domain.Schedule, testExecutor TestExecutor, test domain.Test) error {
	if schedule.Every.Seconds != nil {
		testScheduler.scheduler.Every(uint64(*schedule.Every.Seconds)).Seconds().Do(testExecutor.Execute, test)
	} else if schedule.Every.Minutes != nil {
		testScheduler.scheduler.Every(uint64(*schedule.Every.Minutes)).Minutes().Do(testExecutor.Execute, test)
	} else if schedule.Every.Hours != nil {
		testScheduler.scheduler.Every(uint64(*schedule.Every.Hours)).Hours().Do(testExecutor.Execute, test)
	} else {
		return errors.New("no valid schedule passed")
	}

	return nil
}

func (testScheduler *TestScheduler) Start() {
	<-testScheduler.scheduler.Start()
}
