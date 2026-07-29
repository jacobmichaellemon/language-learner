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
	toLang = "es"

	db, err := download.GetDictionary(fromLang, toLang)

	if err != nil {
		log.Fatal(err)
	}

	column_data, err := data.ReadColumnValues(db, "simple_translation", "written_rep")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(column_data)

	defer db.Close()
}
