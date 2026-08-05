package main

import (
	"database/sql"
	"slices"
	"strings"

	"github.com/jacobmichaellemon/language-learner/internal/data"
)

// language codes -> languages; TODO: commented one exist but do not load properly
var languages = map[string]string{
	//"bg": "Bulgarian",
	"ca": "Catalan",
	"cs": "Czech",
	"da": "Danish",
	"de": "German",
	//"el": "Greek",
	"en": "English",
	"es": "Spanish",
	"fi": "Finnish",
	"fr": "French",
	"ga": "Irish",
	//"id": "Indonesian",
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
	//"ru": "Russian",
	"sv": "Swedish",
	"tr": "Turkish",
	//"zh": "Simplified Chinsese"
}

var specialChars = map[string]string{
	"ca": "à ç é è í ï ó ò ú ü",
	"cs": "á č ď é ě í ň ó ř š ť ú ů ý ž",
	"da": "æ ø å",
	"de": "ä ö ü ß",
	"es": "á é í ó ú ñ",
	"fi": "ä ö å",
	"fr": "â, ê, î, ô, û",
	"ga": "á ḃ ċ ḋ é ḟ ġ í ṁ ó ṗ ṡ ṫ ú",
	"it": "à è é ì ò ó ù",
	"ku": "ç ê î ł ň ř ş û ü",
	"lt": "ą č ė į š ų ū ž",
	"nl": "ë ï ö ü é",
	"no": "æ ø å é ô",
	"pl": "ą ć ę ł ń ó ś ź",
	"pt": "à á â ã é ê í ó ô õ ú ç ü",
	"sv": "å ä ö",
	"tr": "â ç ğ î ı ö ş ü û",
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
