package editor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/GeertJohan/go.rice"
	"github.com/bitrise-io/go-utils/readerutil"
)

const defaultPort = "3000"

// SimpleResponse ...
type SimpleResponse struct {
	Message string `json:"message"`
}

// RegexRequestModel ...
type RegexRequestModel struct {
	Log     string `json:"log"`
	Pattern string `json:"pattern"`
}

// RegexResponseModel ...
type RegexResponseModel struct {
	Matches []string `json:"matches"`
}

func respondWithJSON(w http.ResponseWriter, httpStatusCode int, respModel interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	if err := json.NewEncoder(w).Encode(&respModel); err != nil {
		log.Println(" [!] Exception: RespondWith: Error: ", err)
	}
}

func respondWithErrorMessage(w http.ResponseWriter, format string, v ...interface{}) {
	respondWithJSON(w, 400, SimpleResponse{
		Message: fmt.Sprintf(format, v...),
	})
}

func respondWithSuccessMessage(w http.ResponseWriter, format string, v ...interface{}) {
	respondWithSuccess(w, SimpleResponse{
		Message: fmt.Sprintf(format, v...),
	})
}

func respondWithSuccess(w http.ResponseWriter, obj interface{}) {
	respondWithJSON(w, 200, obj)
}

func testRegexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	log.Println("=> Request")
	if r.Method != "POST" {
		respondWithErrorMessage(w, "Invalid method, only POST is accepted")
		return
	}

	defer r.Body.Close()
	var reqObj RegexRequestModel
	if err := json.NewDecoder(r.Body).Decode(&reqObj); err != nil {
		respondWithErrorMessage(w, "Failed to read JSON input, error: %s", err)
		return
	}

	log.Println(" -> Pattern: ", reqObj.Pattern)
	re, err := regexp.Compile(reqObj.Pattern)
	if err != nil {
		respondWithErrorMessage(w, "Invalid Pattern, error: %s", err)
		return
	}

	log.Println(" -> Log: ", reqObj.Log)
	matches := []string{}
	err = readerutil.WalkLinesString(reqObj.Log, func(line string) error {
		lineMatches := re.FindAllString(line, -1)
		if len(lineMatches) > 0 {
			matches = append(matches, lineMatches...)
		}
		return nil
	})
	if err != nil {
		respondWithErrorMessage(w, "Failed to walk through the log, error: %s", err)
		return
	}

	log.Printf(" -> matches: %+v", matches)

	respondWithJSON(w, 200, RegexResponseModel{Matches: matches})
}

// LaunchEditor ...
func LaunchEditor() error {
	port := envString("PORT", defaultPort)
	log.Printf("Starting web IDE at http://localhost:%s", port)

	http.Handle("/", http.FileServer(rice.MustFindBox("www").HTTPBox()))
	http.HandleFunc("/api/test-regex", testRegexHandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("Can't start HTTP listener: %v", err)
	}
	return nil
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
