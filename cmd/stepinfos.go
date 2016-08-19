// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//

package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// stepinfosCmd represents the stepinfos command
var stepinfosCmd = &cobra.Command{
	Use:   "stepinfos BITRISE-LOG-FILE-PATH",
	Short: "Filter only step infos from the bitrise log",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("stepinfos:")
		if len(args) < 1 {
			return errors.New("No log file path specified")
		}
		logFilePath := args[0]

		return filterStepInfosFromLogFile(logFilePath)
	},
}

func filterStepInfosFromLogFile(logFilePath string) error {
	file, err := os.Open(logFilePath)
	if err != nil {
		return fmt.Errorf("Failed to read log file (%s), error: %s", logFilePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		trimmedLine := strings.TrimSpace(lineStr)
		if len(trimmedLine) > 2 {
			switch trimmedLine[0:2] {
			case "+-", "| ":
				fmt.Println(lineStr)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Failed to scan through the file (%s), error: %s", logFilePath, err)
	}
	return nil
}

func init() {
	RootCmd.AddCommand(stepinfosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stepinfosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stepinfosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
