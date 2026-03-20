package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Definition struct {
	Word       string   `json:"word"`
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
}
type WordInput struct {
	Word string `json:"word"`
}

func main() {

}

func (wd WordInput) getDefinition(w http.ResponseWriter, r *http.Request) string {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error parsing request: %v", err)
		return ""
	}

	searchedWord := r.FormValue("word")
	if searchedWord != "" {
		return ""
	}

	serializedWord := wd{
		Word: searchedWord,
	}

	return ""
}
