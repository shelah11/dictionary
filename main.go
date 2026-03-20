package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type TranslatedResponse struct {
	Contents struct {
		Translated string `json:"translated"`
	} `json:"contents"`
}

type Definition struct {
	Meanings []struct {
		Definitions []struct {
			Definition string `json:"definition"`
		} `json:"definitions"`
	} `json:"meanings"`
}

type SentenceInput struct {
	Sentence    string
	Translation string
	Error       string
	Definitions []string
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, map[string]interface{}{"ActiveTab": "1"})
	})

	http.HandleFunc("/translate", func(w http.ResponseWriter, r *http.Request) {
		var si SentenceInput
		si.getTranslation(w, r)

		tmpl.Execute(w, map[string]interface{}{
			"Data":      si,
			"ActiveTab": "2",
		})
	})

	http.HandleFunc("/definition", func(w http.ResponseWriter, r *http.Request) {
		var si SentenceInput
		si.getDefinition(w, r)

		tmpl.Execute(w, map[string]interface{}{
			"Data":      si,
			"ActiveTab": "1",
		})

	})

	fmt.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Your kept function with logic fixes
func (si *SentenceInput) getTranslation(w http.ResponseWriter, r *http.Request) {
	sentenceInput := r.FormValue("translate-input")

	if sentenceInput == "" {
		si.Error = "Please enter text to translate."
		return
	}

	si.Sentence = sentenceInput
	fmt.Printf("[LOCAL] Translating: \"%s\"\n", si.Sentence)

	words := map[string]string{
		"you":           "thou",
		"your":          "thy",
		"yours":         "thine",
		"are":           "art",
		"do":            "dost",
		"does":          "doth",
		"have":          "hast",
		"will":          "wilt",
		"shall":         "shalt",
		"here":          "hither",
		"there":         "thither",
		"where":         "whither",
		"why":           "wherefore",
		"often":         "oft",
		"before":        "ere",
		"soon":          "anon",
		"morning":       "morrow",
		"listen":        "hark",
		"please":        "prithee",
		"anything":      "aught",
		"nothing":       "nought",
		"hurry":         "hie",
		"notice":        "mark",
		"scoundrel":     "knave",
		"handsome":      "brave",
		"sad":           "heavy",
		"crazy":         "mad",
		"disgusting":    "vile",
		"scold":         "chide",
		"cheat":         "cozen",
		"stop":          "forbear",
		"intelligence":  "wit",
		"always":        "still",
		"from here":     "hence",
		"from where":    "whence",
		"over there":    "yonder",
		"gladly":        "fain",
		"sir":           "sirrah",
		"lord":          "liege",
		"to court":      "woo",
		"created":       "wrought",
		"consider":      "perpend",
		"curse":         "beshrew",
		"bring to life": "quicken",
		"grumpy":        "testy",
		"lowly":         "base",
		"evil":          "ill",
		"perfect":       "absolute",
		"equal":         "egal",
		"dangerously":   "parlous",
	}

	translated := si.Sentence
	for modern, bard := range words {
		translated = strings.ReplaceAll(strings.ToLower(translated), modern, bard)
	}

	si.Translation = strings.Title(translated)
	fmt.Printf("[SUCCESS] Local Result: \"%s\"\n", si.Translation)
}

func (si *SentenceInput) getDefinition(w http.ResponseWriter, r *http.Request) {
	wordInput := r.FormValue("word")

	if wordInput == "" {
		si.Error = "Please enter a word."
		return
	}
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", wordInput)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var definitionResponse []Definition
	err = json.NewDecoder(response.Body).Decode(&definitionResponse)
	if err != nil {
		panic(err)
	}

	si.Definitions = []string{}

	for _, meaning := range definitionResponse[0].Meanings {
		for _, d := range meaning.Definitions {
			si.Definitions = append(si.Definitions, d.Definition)
		}
	}

}
