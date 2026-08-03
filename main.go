package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("assets/*.html"))

func main() {
	app := &QuizApp{
		Tmpl:     template.Must(template.ParseGlob("assets/*.html")),
		Sessions: make(map[string]UserSession),
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Attach struct methods directly as http.HandlerFunc
	http.HandleFunc("/", app.handleCreateQuiz)
	http.HandleFunc("/quiz", app.handleQuiz)
	http.HandleFunc("/submit", app.handleSubmit)
	http.HandleFunc("/results", app.handleResults)
	http.HandleFunc("/reset", app.handleReset)

	log.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
