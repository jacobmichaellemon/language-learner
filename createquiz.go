package main

import (
	"database/sql"
	"fmt"
	"slices"

	"github.com/jacobmichaellemon/language-learner/internal/data"
)

var languages = map[string]string{
	"bg": "Bulgarian",
	"ca": "Catalan",
	"cs": "Czech",
	"da": "Danish",
	"de": "German",
	"el": "Greek",
	"en": "English",
	"es": "Spanish",
	"fi": "Finnish",
	"fr": "French",
	"ga": "Irish",
	"id": "Indonesian",
	"it": "Italian",
	"ja": "Japanese",
	"ku": "Kurdish",
	"la": "Latin",
	"lt": "Lithuanian",
	"mg": "Malagasy",
	"nl": "Dutch",
	"no": "Norwegian",
	"pl": "Polish",
	"pt": "Portuguese",
	"ru": "Russian",
	"sv": "Swedish",
	"tr": "Turkish",
	"zh": "Simplified Chinsese"}

// generate a list of 20 vocab questions to quiz on
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
	return translations, nil
}

// get a valid language code from the list of datasets availible. direction: to/from ex. from english to spanish
func GetValidLanguageCode(direction string) string {
	var language_code string
	for {
		fmt.Printf("Enter a valid language code to translate %s: \n", direction)
		fmt.Scan(&language_code)
		_, ok := languages[language_code]
		if !ok {
			fmt.Println("Invalid language code, please try again!")
			continue
		} else {
			break
		}
	}
	return language_code
}
