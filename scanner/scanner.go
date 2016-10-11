package scanner

import (
	"fmt"
	"io"
	"log"
	"os"
)

//go:generate stringer -type=LogLineType

// LogLineType ...
type LogLineType int

const (
	// BeforeFirstStep ...
	BeforeFirstStep LogLineType = iota
	// StepInfo ...
	StepInfo
	// StepLog ...
	StepLog
	// Summary ...
	Summary
	// AfterSummary ...
	AfterSummary
)

// WalkLogFn ...
type WalkLogFn func(line string, lineType LogLineType)

// WalkLogFile ...
func WalkLogFile(logFilePath string, walkFn WalkLogFn) error {
	file, err := os.Open(logFilePath)
	if err != nil {
		return fmt.Errorf("Failed to read log file (%s), error: %s", logFilePath, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(" [!] Failed to close file, error: ", err)
		}
	}()

	return WalkLog(file, walkFn)
}

// WalkLog ...
func WalkLog(logReader io.Reader, walkFn WalkLogFn) error {
	// patternLineTypeMap := map[string]LogLineType{
	// 	`+------------------------------------------------------------------------------+`: StepInfo,
	// 	`| \([[:digit:]]+)\ .+[[:space:]]+|`:                                               StepInfo,
	// 	`|                                                                              |`: StepInfo,
	// 	`+----+--------------------------------------------------------------+----------+`: StepInfo,
	// }
	// currLogLineType := BeforeFirstStep
	// err := readerutil.WalkLines(logReader, func(line string) error {
	// 	if isMatch, err := regexp.MatchString(pattern, trimmedLine); err != nil {
	// 		return fmt.Errorf("Failed to match line (%s) with regex (%s), error: %s", line, pattern, err)
	// 	} else if isMatch {
	// 		fmt.Println(line)
	// 	}

	// 	walkFn(line, currLogLineType)
	// 	return nil
	// })
	// if err != nil {
	// 	return fmt.Errorf("Failed to scan through the log, error: %s", err)
	// }
	return nil
}
