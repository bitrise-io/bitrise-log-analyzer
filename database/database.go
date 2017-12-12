package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/bitrise-tools/bitrise-log-analyzer/build"
)

const DBUrl = "https://gist.githubusercontent.com/kdobmayer/bea24c0d9b4cd9d6ee5011a4df11e5c8/raw/bea631ad937f0d858c0197036d9e059de577c443/log_analyzer.json"

type DataBase struct {
	Data []Item `json:"data"`
}

type Item struct {
	Pattern string `json:"pattern"`
	Answer  string `json:"answer"`
}

func (db DataBase) Search(step build.Step) (string, error) {
	for _, item := range db.Data {
		r, err := regexp.Compile(item.Pattern)
		if err != nil {
			fmt.Printf("can't compile regexp (%s): %v\n", item.Pattern, err)
			continue
		}
		for _, line := range step.Lines {
			if r.MatchString(line) {
				return item.Answer, nil
			}
		}
	}
	return "", errors.New("no matches found")
}

func New(url string) (DataBase, error) {
	var db DataBase
	if err := download(url, &db); err != nil {
		return DataBase{}, err
	}
	return db, nil
}

func download(url string, db *DataBase) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download url %q: %v", url, err)
	}
	defer func() {
		if cerr := resp.Body.Close(); err == nil {
			err = cerr
		}
	}()
	return json.NewDecoder(resp.Body).Decode(db)
}
