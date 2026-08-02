package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/jacobmichaellemon/language-learner/internal/download"
)

var tmpl = template.Must(template.ParseGlob("assets/*.html"))

func main() {

	/*fmt.Println("Welcome to language learner!")
	fmt.Println("The availible languages are: ")
	PrintLanguages()*/

	fromLang := GetValidLanguageCode(From)
	toLang := GetValidLanguageCode(To)

	db, err := download.GetDictionary(fromLang, toLang)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	questions, err := GetQuizWords(db, 3.0)
	if err != nil {
		log.Fatal(err)
	}

	app := &QuizApp{
		Tmpl:      template.Must(template.ParseGlob("assets/*.html")),
		Sessions:  make(map[string]UserSession),
		DB:        db,
		Questions: questions,
	}

	// 1. Create a file server pointing to your local 'static' directory
	fs := http.FileServer(http.Dir("static"))

	// 2. Strip "/static/" from the request path before passing it to the file server
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Attach struct methods directly as http.HandlerFunc
	http.HandleFunc("/", app.handleQuiz)
	http.HandleFunc("/submit", app.handleSubmit)
	http.HandleFunc("/results", app.handleResults)
	http.HandleFunc("/reset", app.handleReset)

	http.ListenAndServe(":8080", nil)

	/**score := StartQuiz(questions, toLang, fromLang)

	fmt.Printf("Final Score: %d\n", score)**/

}
