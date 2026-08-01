package main

import (
	"fmt"
	"log"

	"github.com/jacobmichaellemon/language-learner/internal/download"
)

func main() {

	fmt.Println("Welcome to language learner!")
	fmt.Println("The availible languages are: ")
	for key, value := range languages {
		fmt.Printf("Language Code: %s Language: %s\n", key, value)
	}

	fromLang := GetValidLanguageCode("from")
	toLang := GetValidLanguageCode("to")

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
