package pattern

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/bitrise-io/go-utils/readerutil"
)

// Model ...
type Model struct {
	Line string
}

// Matcher ...
type Matcher struct {
	// patterns - should NOT be modified once assigned (through `NewMatcher`)!
	patterns []Model
	// matches - index of `patterns` which matched
	matches map[int]bool
}

// NewMatcher ...
func NewMatcher(ptrns []Model) Matcher {
	return Matcher{
		patterns: ptrns,
		matches:  map[int]bool{},
	}
}

// Results ...
func (matcher *Matcher) Results() []Model {
	results := []Model{}

	for idx := range matcher.matches {
		results = append(results, matcher.patterns[idx])
	}

	return results
}

// ProcessLine - process a single line
func (matcher *Matcher) ProcessLine(line string) error {
	for idx, aPattern := range matcher.patterns {
		if _, isFound := matcher.matches[idx]; isFound {
			// skip, already matched
			continue
		}
		r, err := regexp.Compile(aPattern.Line)
		if err != nil {
			return fmt.Errorf("Invalid pattern: %s | error: %s", aPattern, err)
		}
		if r.MatchString(line) {
			matcher.matches[idx] = true
		}
	}
	return nil
}

// ProcessText - process (potentially multi line) text
func (matcher *Matcher) ProcessText(text string) error {
	return matcher.ProcessReader(strings.NewReader(text))
}

// ProcessReader - process the reader, until EOF
func (matcher *Matcher) ProcessReader(textReader io.Reader) error {
	err := readerutil.WalkLines(textReader, func(line string) error {
		return matcher.ProcessLine(line)
	})

	if err != nil {
		return fmt.Errorf("Failed to process reader, error: %s", err)
	}

	return nil
}
