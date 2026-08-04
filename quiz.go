package main

import (
	"database/sql"
	"slices"
	"strings"

	"github.com/jacobmichaellemon/language-learner/internal/data"
)

type Direction string

// string mapping
const (
	From Direction = "from"
	To   Direction = "to"
)

var languages = map[string]string{
	//"bg": "Bulgarian",
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
	//"ja": "Japanese",
	"ku": "Kurdish",
	//"la": "Latin",
	"lt": "Lithuanian",
	"mg": "Malagasy",
	"nl": "Dutch",
	"no": "Norwegian",
	"pl": "Polish",
	"pt": "Portuguese",
	"ru": "Russian",
	"sv": "Swedish",
	"tr": "Turkish",
	//"zh": "Simplified Chinsese"
}

var specialChars = map[string]string{
	"es": "á é í ó ú ñ",
}

// generate a list of X vocab questions to quiz on based on the importance of the words
func GetQuizWords(db *sql.DB, numQuestions int, importance float32) ([]data.Translation, error) {
	var translations []data.Translation
	var words []string
	for len(translations) < numQuestions {
		newTranslations, err := data.GetRandomTranslations(db, importance)
		newTranslation := data.GetRandomWord(newTranslations)
		if slices.Contains(words, strings.ToLower(newTranslation.Written_Rep)) || newTranslation.Written_Rep == "" {
			continue
		}
		words = append(words, strings.ToLower(newTranslation.Written_Rep))
		translations = append(translations, newTranslation)
		if err != nil {
			return nil, err
		}
	}
	return translations, nil
}
