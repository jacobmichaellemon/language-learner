package main

import (
	"log"

	"github.com/jacobmichaellemon/language-learner/internal/download"
)

func main() {

	var fromLang, toLang string
	fromLang = "en"
	toLang = "de"

	db, err := download.GetDictionary(fromLang, toLang)

	if err != nil {
		log.Fatal(err)
	}

	_, err = CreateQuiz(db)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
