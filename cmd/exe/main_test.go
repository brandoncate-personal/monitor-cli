package main

import (
	"os"
	"testing"
)

func Test_Main_Success(t *testing.T) {
	os.Args = []string{"cmd", "--yaml=data/default_test_data.yaml"}

	main()
}
