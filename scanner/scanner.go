package scanner

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/bitrise-io/go-utils/readerutil"
)

//go:generate stringer -type=LogLineType

// LogLineType ...
type LogLineType int

const (
	// BeforeFirstStep ...
	BeforeFirstStep LogLineType = iota
	// StepInfoHeader ...
	StepInfoHeader
	// StepLog ...
	StepLog
	// StepInfoFooter ...
	StepInfoFooter
	// BetweenSteps ...
	BetweenSteps
	// BuildSummary ...
	BuildSummary
	// StepInfoHeaderOrBuildSummarySectionStarter ...
	StepInfoHeaderOrBuildSummarySectionStarter
	// AfterBuildSummary ...
	AfterBuildSummary
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

func regexListMatch(text string, regexs []*regexp.Regexp) bool {
	for _, aRegex := range regexs {
		if isMatch := aRegex.MatchString(text); isMatch {
			return true
		}
	}
	return false
}

// WalkLog ...
func WalkLog(logReader io.Reader, walkFn WalkLogFn) error {
	stepInfoHeaderOrBuildSummarySectionStarterPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\+------------------------------------------------------------------------------\+`),
	}
	stepHeaderIndicatorPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\| \([[:digit:]]+\) .+[[:space:]]+\|`),
		regexp.MustCompile(`\|                                                                              \|`),
	}
	stepHeaderAdditionalPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\| .* \|`),
		regexp.MustCompile(`\+------------------------------------------------------------------------------\+`),
	}
	stepFooterIndicatorPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\|                                                                              \|`),
		regexp.MustCompile(`\+---\+---------------------------------------------------------------\+----------\+`),
		regexp.MustCompile(`\|.+\|.+[[:space:]]+\| [0-9\.]+ sec[[:space:]]+\|`),
	}
	buildSummaryIndicatorPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\|                               bitrise summary                                \|`),
	}
	buildSummaryAdditionalPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\| .* \|`),
		regexp.MustCompile(`\+---\+---------------------------------------------------------------\+----------\+`),
		regexp.MustCompile(`\|    \| title                                                        \| time \(s\) \|`),
		regexp.MustCompile(`\+------------------------------------------------------------------------------\+`),
	}

	currLogLineType := BeforeFirstStep
	err := readerutil.WalkLines(logReader, func(line string) error {
		switch currLogLineType {
		case BeforeFirstStep:
			if regexListMatch(line, stepInfoHeaderOrBuildSummarySectionStarterPatterns) {
				currLogLineType = StepInfoHeaderOrBuildSummarySectionStarter
			} else if regexListMatch(line, stepHeaderIndicatorPatterns) {
				currLogLineType = StepInfoHeader
			}
		case StepInfoHeader:
			if !regexListMatch(line, stepHeaderIndicatorPatterns) &&
				!regexListMatch(line, stepHeaderAdditionalPatterns) {
				currLogLineType = StepLog
			}
		case StepLog:
			if regexListMatch(line, stepFooterIndicatorPatterns) {
				currLogLineType = StepInfoFooter
			}
			if regexListMatch(line, stepInfoHeaderOrBuildSummarySectionStarterPatterns) {
				currLogLineType = StepInfoHeaderOrBuildSummarySectionStarter
			}
		case StepInfoFooter:
			if !regexListMatch(line, stepFooterIndicatorPatterns) {
				currLogLineType = BetweenSteps
			}
		case BetweenSteps:
			if regexListMatch(line, stepInfoHeaderOrBuildSummarySectionStarterPatterns) {
				currLogLineType = StepInfoHeaderOrBuildSummarySectionStarter
			} else if regexListMatch(line, stepHeaderIndicatorPatterns) {
				currLogLineType = StepInfoHeader
			}
		case StepInfoHeaderOrBuildSummarySectionStarter:
			if regexListMatch(line, buildSummaryIndicatorPatterns) {
				currLogLineType = BuildSummary
			} else if regexListMatch(line, stepHeaderIndicatorPatterns) {
				currLogLineType = StepInfoHeader
			}
		case BuildSummary:
			if !regexListMatch(line, buildSummaryAdditionalPatterns) {
				currLogLineType = AfterBuildSummary
			}
		case AfterBuildSummary:
			if regexListMatch(line, stepFooterIndicatorPatterns) {
				currLogLineType = StepInfoFooter
			}
		}

		walkFn(line, currLogLineType)
		return nil
	})
	if err != nil {
		return fmt.Errorf("Failed to scan through the log, error: %s", err)
	}
	return nil
}
