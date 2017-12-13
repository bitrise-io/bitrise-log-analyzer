package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/bitrise-tools/bitrise-log-analyzer/scanner"
	"github.com/spf13/cobra"
)

var raw *bool

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a log",
	Long:  `Scan a log`,
	Args:  cobra.MinimumNArgs(1),
	RunE:  scan,
}

func init() {
	RootCmd.AddCommand(scanCmd)
	raw = scanCmd.Flags().Bool("raw", false, "raw log")
}

func scan(cmd *cobra.Command, args []string) error {
	logFilePath := args[0]

	db, err := scanner.NewDatabase()
	if err != nil {
		return err
	}

	if *raw {
		content, err := ioutil.ReadFile(logFilePath)
		if err != nil {
			return err
		}
		lines := strings.Split(string(content), "\n")
		answer, err := db.Search(lines)
		if err != nil {
			return fmt.Errorf("can't find possible solution: %v", err)
		}
		fmt.Printf("\nPossible solution:\n%s\n", answer)
		return nil
	}

	steps := []scanner.Step{}
	scanned := true
	var currentStep *scanner.Step

	if err := scanner.WalkLogFile(logFilePath, func(line string, lineType scanner.LogLineType) {
		switch lineType {
		case scanner.StepInfoHeader:
			if scanned {
				scanned = false
				currentStep = new(scanner.Step)
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
	}); err != nil {
		return err
	}

	return search(db, steps)
}

func search(db scanner.Database, steps []scanner.Step) error {
	for _, step := range steps {
		if step.Status == scanner.Failed {
			fmt.Println("Failed step:")
			fmt.Printf("- id: %s\n- duration: %v\n\nLog:\n%s\n",
				step.ID, step.Duration, strings.Join(step.Lines, "\n"))

			answer, err := db.Search(step.Lines)
			if err != nil {
				return fmt.Errorf("can't find possible solution: %v", err)
			}
			fmt.Printf("\nPossible solution:\n%s\n", answer)
			return nil
		}
	}
	return errors.New("There was no failed step")
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

func status(line string) scanner.Status {
	switch {
	case strings.Contains(line, string(scanner.Red)):
		return scanner.Failed
	case strings.Contains(line, string(scanner.Green)):
		return scanner.Success
	default:
		return scanner.Skipped
	}
}
