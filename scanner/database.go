package scanner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bitrise-io/go-utils/command/git"
)

const repo = "https://github.com/bitrise-tools/bitrise-log-analyzer-patterns.git"

var dir = filepath.Join(os.Getenv("HOME"), ".bitrise-log-analyzer")

type Item struct {
	Pattern string `json:"pattern"`
	Answer  string `json:"answer"`
}

type Database struct {
	Data []Item `json:"data"`
}

func NewDatabase() (Database, error) {
	var db Database
	if err := initRepo(); err != nil {
		return Database{}, err
	}

	path := filepath.Join(dir, "patterns.json")
	file, err := os.Open(path)
	if err != nil {
		return Database{}, err
	}
	defer func() {
		if cerr := file.Close(); err == nil {
			err = cerr
		}
	}()

	if err := json.NewDecoder(file).Decode(&db); err != nil {
		return Database{}, err
	}
	return db, nil
}

func (db Database) searchInLines(lines []string) (string, error) {
	for _, item := range db.Data {
		r, err := regexp.Compile(item.Pattern)
		if err != nil {
			fmt.Printf("can't compile regexp (%s): %v\n", item.Pattern, err)
			continue
		}
		for _, line := range lines {
			if r.MatchString(line) {
				return item.Answer, nil
			}
		}
	}
	return "", errors.New("no matches found")
}

func (db Database) SearchInRawLog(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	answer, err := db.searchInLines(lines)
	if err != nil {
		return fmt.Errorf("can't find possible solution: %v", err)
	}
	fmt.Printf("\nPossible solution:\n%s\n", answer)
	return nil
}

func (db Database) SearchInSteps(steps []Step) error {
	for _, step := range steps {
		if step.Status == Failed {
			fmt.Println("Failed step:")
			fmt.Printf("- id: %s\n- duration: %v\n\nLog:\n%s\n",
				step.ID, step.Duration, strings.Join(step.Lines, "\n"))

			answer, err := db.searchInLines(step.Lines)
			if err != nil {
				return fmt.Errorf("can't find possible solution: %v", err)
			}
			fmt.Printf("\nPossible solution:\n%s\n", answer)
			return nil
		}
	}
	return errors.New("There was no failed step")
}

func initRepo() error {
	var g git.Git
	g, err := git.New(dir)
	if err != nil {
		return err
	}

	if file, err := os.Stat(filepath.Join(dir, ".git")); err == nil && file.IsDir() {
		return g.Pull().SetStdout(os.Stdout).SetStderr(os.Stderr).Run()
	}

	return g.Clone(repo).SetStdout(os.Stdout).SetStderr(os.Stderr).Run()
}
