package main

import (
	"fmt"
	"os"

	"github.com/bitrise-tools/bitrise-log-analyzer/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
