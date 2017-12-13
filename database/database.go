package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/bitrise-io/go-utils/command/git"
)

const repo = "https://github.com/bitrise-tools/bitrise-log-analyzer-patterns.git"

var dir = filepath.Join(os.Getenv("HOME"), ".bitrise-log-analyzer")

type DataBase struct {
	Data []Item `json:"data"`
}

type Item struct {
	Pattern string `json:"pattern"`
	Answer  string `json:"answer"`
}

func (db DataBase) Search(lines []string) (string, error) {
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

func New() (DataBase, error) {
	var db DataBase
	if err := initRepo(); err != nil {
		return DataBase{}, err
	}

	path := filepath.Join(dir, "patterns.json")
	file, err := os.Open(path)
	if err != nil {
		return DataBase{}, err
	}
	defer func() {
		if cerr := file.Close(); err == nil {
			err = cerr
		}
	}()

	if err := json.NewDecoder(file).Decode(&db); err != nil {
		return DataBase{}, err
	}
	return db, nil
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
