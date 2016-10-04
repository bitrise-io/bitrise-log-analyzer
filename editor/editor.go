package editor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "3000"

// RegexResponse ...
type RegexResponse struct {
	Message string `json:"message"`
}

func respondWith(w http.ResponseWriter, httpStatusCode int, respModel interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	if err := json.NewEncoder(w).Encode(&respModel); err != nil {
		log.Println(" [!] Exception: RespondWith: Error: ", err)
	}
}

func testRegexHandler(w http.ResponseWriter, r *http.Request) {
	resp := RegexResponse{
		Message: "test",
	}

	respondWith(w, 200, resp)
}

// LaunchEditor ...
func LaunchEditor() error {
	port := envString("PORT", defaultPort)
	log.Printf("Starting web IDE at http://localhost:%s", port)

	// _, filename, _, _ := runtime.Caller(1)
	// fs := http.FileServer(http.Dir(path.Join(filepath.Dir(filename), "www")))

	// http.Handle("/", fs)
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
