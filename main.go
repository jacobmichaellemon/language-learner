package main

import (
	"fmt"
	"log"

	"github.com/jacobmichaellemon/language-learner/internal/download"
)

func main() {

	fmt.Println("Welcome to language learner!")
	fmt.Println("The availible languages are: ")
	PrintLanguages()

	fromLang := GetValidLanguageCode(From)
	toLang := GetValidLanguageCode(To)

	db, err := download.GetDictionary(fromLang, toLang)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	translations, err := GetQuizWords(db, 3.0)
	if err != nil {
		log.Fatal(err)
	}

	score := StartQuiz(translations, toLang, fromLang)

	fmt.Printf("Final Score: %d\n", score)

}
