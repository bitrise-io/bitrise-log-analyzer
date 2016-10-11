package pattern

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/bitrise-io/go-utils/readerutil"
)

// Model ...
type Model struct {
	Lines []string
}

// Matcher ...
type Matcher struct {
	// patterns - should NOT be modified once assigned (through `NewMatcher`)!
	patterns []Model
	// patternLineMatchLookup - keeps track of which pattern lines were matched/should be checked next.
	// The value of this map if the index of the `.Lines` of the related `pattern[idx]`
	//  which should be checked for a match
	// map is: index of `patterns` => pattern.Lines index which should match with the next line
	//  in order to declare the pattern as a match
	patternLineMatchLookup map[int]int
	// matches - index of `patterns` which matched
	matches map[int]bool
}

// NewMatcher ...
func NewMatcher(ptrns []Model) Matcher {
	filteredPatterns := []Model{}
	filteredPatternLookupMap := map[int]int{}

	filteredPatternsIdx := 0
	for _, aPattern := range ptrns {
		if len(aPattern.Lines) < 1 {
			log.Println("Skipping empty pattern (no Lines defined)")
			continue
		}
		filteredPatterns = append(filteredPatterns, aPattern)
		// init to pattern.Lines idx 0
		filteredPatternLookupMap[filteredPatternsIdx] = 0
		filteredPatternsIdx++
	}

	matcher := Matcher{
		patterns:               filteredPatterns,
		patternLineMatchLookup: filteredPatternLookupMap,
		matches:                map[int]bool{},
	}

	return matcher
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

		currPatternLineToMatchIdx := matcher.patternLineMatchLookup[idx]
		patternLineToMatch := aPattern.Lines[currPatternLineToMatchIdx]
		r, err := regexp.Compile(patternLineToMatch)
		if err != nil {
			return fmt.Errorf("Invalid pattern: %s | error: %s", aPattern, err)
		}

		if r.MatchString(line) {
			// line matched, check if this was the last in the pattern's Lines array/slice
			if len(aPattern.Lines)-1 <= currPatternLineToMatchIdx {
				// if it was, then we just had a match! ;)
				matcher.matches[idx] = true
			} else {
				// otherwise bump the pattern's .Lines index for the next check
				matcher.patternLineMatchLookup[idx] = currPatternLineToMatchIdx + 1
			}
		} else {
			// line did not match - reset the pattern.Lines index to 0
			matcher.patternLineMatchLookup[idx] = 0
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
