package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("GET /{word}", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /{word} request received")
		word := r.PathValue("word")
		if word == "" {
			http.Error(w, "word can't be empty", http.StatusBadRequest)
			return
		}

		word = strings.ToLower(word)
		rg := regexp.MustCompile(`^[a-z]+$`)
		if !rg.MatchString(word) {
			http.Error(w, "word must contain only letters", http.StatusBadRequest)
			return
		}

		url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("couldn't fetch word from dictionary, error: %v", err)
			http.Error(w, "couldn't fetch word from dictionary", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if ct := resp.Header.Get("Content-Type"); ct != "" {
			w.Header().Add("Content-Type", ct)
		}
		w.WriteHeader(resp.StatusCode)

		log.Printf("got response from dictionary for the word: \"%s\"\n", word)
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, "couldn't get definition of the word", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /health request received")
		resp := struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		}
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("couldn't encode health response, err: %v", err)
			return
		}
	})

	log.Println("starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error while running server %v", err)
	}
}
