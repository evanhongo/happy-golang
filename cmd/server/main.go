package main

import (
	"os"

	"github.com/evanhongo/happy-golang/pkg/logger"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "Startup happy-golang",
	Short: "Startup happy-golang",
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := CreateServer()
		if err != nil {
			return err
		}
		if err := server.Start(); err != nil {
			return err
		}

		return nil
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
