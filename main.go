package main

import (
	"os"

	"github.com/WuerthIT/errlogbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
