package cmd

import (
	"errors"
	"fmt"

	"github.com/bitrise-tools/bitrise-log-analyzer/scanner"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a log",
	Long:  `Scan a log`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("No log file specified")
		}
		logFilePath := args[0]
		return scanner.WalkLogFile(logFilePath, func(line string, lineType scanner.LogLineType) {
			fmt.Printf("[%s] %s\n", lineType, line)
		})
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)
}
