package services

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
)

type GoTestRunner struct{}

// cases:
// 1) test passes
// 2) test fails
// 3) test error (build, compile)
func (testRunner GoTestRunner) Run(test domain.Test) (domain.TestResults, error) {
	goExe, err := exec.LookPath("go")
	if err != nil {
		log.Fatal("failed to find go executable: ", err)
	}

	// build args from params
	args := []string{goExe, "test", "--json"}
	if test.Name != "" {
		args = append(args, "-run", test.Name)
	} else {
		args = append(args, ".")
	}
	// never cache results
	args = append(args, "-count", "1")

	var stdout, stderr bytes.Buffer
	testCmd := &exec.Cmd{
		Dir:    test.Dir,
		Path:   goExe,
		Args:   args,
		Stdout: &stdout,
		Stderr: &stderr,
	}

	err = testCmd.Run()

	if err == nil {
		// Test ran without any problems, return json
		return domain.TestResults{
			Name:       test.Name,
			StatusCode: domain.OK,
			Stdout:     stdout.String(),
			Stderr:     stderr.String(),
		}, nil
	}

	exitError, ok := err.(*exec.ExitError)
	if !ok {
		log.Fatalf("error: '%s' failed with non exit error %s",
			testCmd.String(), err,
		)
	}

	switch exc := exitError.ExitCode(); exc {
	case 1:
		// `go test` returns 1 when tests fail, this is fine
		return domain.TestResults{
			Name:       test.Name,
			StatusCode: domain.ERROR,
			Stdout:     stdout.String(),
			Stderr:     stderr.String(),
		}, nil
	case 2:
		//  go test returns 2 on a compilation / build error
		log.Printf("'%s' returned exit code %d: %s",
			testCmd.String(), exc, err,
		)
		return domain.TestResults{
			Name:       test.Name,
			StatusCode: domain.ERROR,
			Stdout:     stdout.String(),
			Stderr:     stderr.String(),
		}, err
	default:
		log.Fatalf("error: '%s' failed with exit error %d: %s",
			testCmd.String(), exc, err,
		)
	}
	return domain.TestResults{
		Name:       test.Name,
		StatusCode: domain.INTERNAL_ERROR,
		Stdout:     stdout.String(),
		Stderr:     stderr.String(),
	}, err
}
