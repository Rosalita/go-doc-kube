package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title    string
	Answered int
	Correct  int
	Question Question
}

type Question struct {
	Qsn          string
	Ans          []string
	CorrectIndex int
}

func main() {
	http.HandleFunc("/", Quiz)
	http.HandleFunc("/score", Score)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Quiz displays a question and four possible answers
func Quiz(w http.ResponseWriter, r *http.Request) {

	page := Page{
		Title:    "Rosie's Gopher Quiz",
		Answered: 0,
		Correct:  0,
		Question: Question{
			Qsn:          "Where do Gopher's live?",
			Ans:          []string{"In tall trees", "Under the sea", "Underground", "In Rainforests"},
			CorrectIndex: 2,
		},
	}

	t, err := template.ParseFiles("quiz.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	if err := t.Execute(w, page); err != nil {
		log.Print("template executing error: ", err)
	}
}

// Score scores a quiz answer
func Score(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	answer := r.Form.Get("quizSelect")
	fmt.Println(answer)
}
