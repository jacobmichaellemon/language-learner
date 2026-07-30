package main

import (
	"fmt"
	"log"

	"github.com/jacobmichaellemon/language-learner/internal/data"
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

	translations, err := data.GetTranslations(db, 5.0)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range translations {
		fmt.Println(value)
	}

	defer db.Close()
}
