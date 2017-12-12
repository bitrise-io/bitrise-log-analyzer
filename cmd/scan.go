package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bitrise-tools/bitrise-log-analyzer/build"
	"github.com/bitrise-tools/bitrise-log-analyzer/database"
	"github.com/bitrise-tools/bitrise-log-analyzer/scanner"
	"github.com/spf13/cobra"
)

var steps = []build.Step{}
var db database.DataBase

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:               "scan",
	Short:             "Scan a log",
	Long:              `Scan a log`,
	Args:              cobra.MinimumNArgs(1),
	PersistentPreRunE: initDB,
	PreRunE:           scan,
	RunE:              search,
}

func init() {
	RootCmd.AddCommand(scanCmd)
}

func parseHeader(line, field string) string {
	stepID := strings.Split(line, "|")[1]
	stepID = strings.TrimSpace(stepID)
	return strings.TrimPrefix(stepID, field+": ")
}

func duration(line string) (time.Duration, error) {
	durationAsString := strings.Split(line, "|")[3]
	durationAsString = strings.Split(durationAsString, " ")[1]
	durationAsFloat, err := strconv.ParseFloat(durationAsString, 32)
	if err != nil {
		return 0, err
	}
	return time.Duration(durationAsFloat*1000) * time.Millisecond, nil
}

func status(line string) build.Status {
	switch {
	case strings.Contains(line, string(build.Red)):
		return build.Failed
	case strings.Contains(line, string(build.Green)):
		return build.Success
	default:
		return build.Skipped
	}
}

func scan(cmd *cobra.Command, args []string) error {
	logFilePath := args[0]
	scanned := true
	var currentStep *build.Step

	return scanner.WalkLogFile(logFilePath, func(line string, lineType scanner.LogLineType) {
		switch lineType {
		case scanner.StepInfoHeader:
			if scanned {
				scanned = false
				currentStep = new(build.Step)
			} else if strings.HasPrefix(line, "| id:") {
				currentStep.ID = parseHeader(line, "id")
			} else if strings.HasPrefix(line, "| version:") {
				currentStep.Version = parseHeader(line, "version")
			}
			return
		case scanner.StepLog:
			currentStep.Lines = append(currentStep.Lines, line)
		case scanner.StepInfoFooter:
			if strings.HasPrefix(line, "| \x1b[") {
				currentStep.Status = status(line)
				dur, err := duration(line)
				if err != nil {
					fmt.Printf("Can't get duration from %s, error: %v\n", line, err)
				}
				currentStep.Duration = dur
				steps = append(steps, *currentStep)
			}
		}
		scanned = true
	})
}

func search(cmd *cobra.Command, args []string) error {
	for _, step := range steps {
		if step.Status == build.Failed {
			fmt.Println("Failed step:")
			fmt.Printf("- id: %s\n- duration: %v\n\nLog:\n%s\n",
				step.ID, step.Duration, strings.Join(step.Lines, "\n"))

			answer, err := db.Search(step)
			if err != nil {
				return fmt.Errorf("can't find possible solution: %v", err)
			}
			fmt.Printf("\nPossible solution:\n%s\n", answer)
			return nil
		}
	}
	return errors.New("There was no failed step")
}

func initDB(cmd *cobra.Command, args []string) error {
	var err error
	db, err = database.New(database.DBUrl)
	if err != nil {
		return fmt.Errorf("can't create db: %v", err)
	}
	return nil
}
