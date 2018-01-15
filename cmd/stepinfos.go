package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bitrise-io/go-utils/readerutil"
	"github.com/spf13/cobra"
)

const stepInfoPattern = `(?i)^\| .+\| [0-9.]+ sec[[:space:]]*\|$`

var onlyTimes *bool

// stepinfosCmd represents the stepinfos command
var stepinfosCmd = &cobra.Command{
	Use:   "stepinfos BITRISE-LOG-FILE-PATH",
	Short: "Filter only step infos from the bitrise log",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("stepinfos:")
		logFilePath := args[0]
		return filterStepInfosFromLogFile(logFilePath, *onlyTimes)
	},
}

func filterStepInfosFromLogFile(logFilePath string, onlyTimes bool) error {
	file, err := os.Open(logFilePath)
	if err != nil {
		return fmt.Errorf("Failed to read log file (%s), error: %s", logFilePath, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(" [!] Failed to close file, error: ", err)
		}
	}()

	re := regexp.MustCompile(stepInfoPattern)
	return readerutil.WalkLines(file, func(line string) error {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "+-") || strings.HasPrefix(line, "| ") {
			if !onlyTimes || re.MatchString(line) {
				fmt.Println(line)
			}
		}
		return nil
	})
}

func init() {
	RootCmd.AddCommand(stepinfosCmd)
	onlyTimes = stepinfosCmd.Flags().Bool("only-times", false, "If enabled will only print step run times")
}
