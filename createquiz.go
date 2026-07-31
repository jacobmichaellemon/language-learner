package main

import (
	"database/sql"
	"fmt"
	"slices"

	"github.com/jacobmichaellemon/language-learner/internal/data"
)

func CreateQuiz(db *sql.DB) ([]data.Translation, error) {
	var translations []data.Translation
	for len(translations) < 20 {
		new_translation, err := data.GetRandomTranslation(db, 3.5)
		if slices.Contains(translations, new_translation) || new_translation.Written_Rep == "" {
			continue
		}
		translations = append(translations, new_translation)
		if err != nil {
			return nil, err
		}
	}
	for _, value := range translations {
		//fmt.Printf("lexentry: %s\n", value.Lexentry.String)
		fmt.Printf("written representation: %s\n", value.Written_Rep)
		fmt.Printf("translation list: %s\n", value.TransList)
		fmt.Println("===========")
	}
	return translations, nil
}
