package main

import (
	"database/sql"
	"fmt"
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

var specialChars = map[string]string{
	"es": "á é í ó ú ñ",
}

// generate a list of 20 vocab questions to quiz on
func GetQuizWords(db *sql.DB, importance float32) ([]data.Translation, error) {
	var translations []data.Translation
	for len(translations) < 3 {
		newTranslations, err := data.GetRandomTranslations(db, importance)
		newTranslation := data.GetRandomWord(newTranslations)
		if slices.Contains(translations, newTranslation) || newTranslation.Written_Rep == "" {
			continue
		}
		translations = append(translations, newTranslation)
		if err != nil {
			return nil, err
		}
	}
	return translations, nil
}

// generate a list of 20 vocab questions to quiz on
func StartQuiz(translations []data.Translation, toLang string, fromLang string) int {
	var score int
	toLanguage := languages[toLang]
	fromLanguage := languages[fromLang]
	fmt.Printf("Quiz starting: guess the %s from the %s word given! Good luck!\n", toLanguage, fromLanguage)
	for i, translation := range translations {
		fmt.Printf("Question #%d: \n", i+1)
		fmt.Printf("%s - %s: ", translation.Written_Rep, translation.Sense.String)
		var guess string
		isCorrect := false
		fmt.Scan(&guess)
		if guess == "!" {
			fmt.Printf("%s\n", specialChars[toLang])
			fmt.Scan(&guess)
		}
		for word := range strings.SplitSeq(translation.TransList, "|") {
			if strings.EqualFold(guess, strings.TrimSpace(word)) {
				fmt.Println("Correct!")
				isCorrect = true
				score++
				break
			}
		}
		if !isCorrect {
			fmt.Println("Wrong")
		}
	}
	return score
}

// get a valid language code from the list of datasets availible. direction: to/from ex. from english to spanish
func GetValidLanguageCode(direction Direction) string {
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

// lists lanaguages availible and their language codes
func PrintLanguages() {
	for key, value := range languages {
		fmt.Printf("Language Code: %s Language: %s\n", key, value)
	}
}
