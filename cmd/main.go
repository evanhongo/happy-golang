package main

import (
	"os"

	"github.com/evanhongo/happy-golang/pkg/logger"
)

func main() {
	cmd, err := CreateCmd()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
