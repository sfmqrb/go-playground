package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/jdkato/prose"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		userInput := string(body)
		response := getResponse(userInput)
		w.Write([]byte(response))
	}
}

func getResponse(userInput string) string {
	// Use NLP to extract information from user input
	doc, _ := prose.NewDocument(userInput)
	for _, ent := range doc.Entities() {
		if ent.Label == "GPE" {
			return fmt.Sprintf("You're asking about %s. I'm sorry, I don't know much about that place.", ent.Text)
		}
	}
	if doc.Sents()[0].Text == "How are you?" {
		return "I'm doing well, thanks for asking!"
	} else {
		return "I'm sorry, I don't understand. Can you please ask me something else?"
	}
}

