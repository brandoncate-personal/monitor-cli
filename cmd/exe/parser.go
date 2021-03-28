package main

import (
	"flag"
)

type Arguments struct {
	yaml string
}

func ParseArgs() Arguments {
	yaml := flag.String("yaml", "", "Path to yaml file to parse.")
	flag.Parse()

	return Arguments{
		yaml: *yaml,
	}
}
