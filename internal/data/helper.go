package data

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"strings"

	_ "modernc.org/sqlite"
)

type Translation struct {
	Lexentry    sql.NullString
	Sense_Num   sql.NullString
	Sense       sql.NullString
	Written_Rep string
	TransList   string
	Score       float64
	Is_Good     int
	Imporatance float64
}

func GetRandomWord(translations []Translation) Translation {
	length := len(translations)
	if length == 0 {
		return Translation{}
	}
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
	return translations[randomInt.Int64()]
}

func getRandomChar() string {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
	return letters[randomInt.Int64()]
}

// ReadColumnValues fetches all values from a specific column in a table.
func ReadColumnValues(db *sql.DB, tableName, columnName string, importance float32) ([]string, error) {
	// Construct the query. (Use parameter substitution for values, but table/column
	// names in SQL must be formatted into the query string).
	query := fmt.Sprintf("SELECT %s FROM %s WHERE importance >= %f", columnName, tableName, importance)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query column: %w", err)
	}
	// Always close rows to free database connections and memory
	defer rows.Close()

	var results []string

	// Iterate through all returned rows
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, value)
	}

	// Check if any error occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return results, nil
}

func GetRandomTranslation(db *sql.DB, importance float32) (Translation, error) {
	// Query all 7 columns from the translation table
	randomChar := getRandomChar()
	stringRandomWildcard := strings.Join([]string{randomChar, "%"}, "")
	query := fmt.Sprintf("SELECT * FROM translation WHERE written_rep LIKE '%s' AND importance >= %f LIMIT 1", stringRandomWildcard, importance)

	rows, err := db.Query(query)
	if err != nil {
		return Translation{}, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var translation Translation

	for rows.Next() {

		// Pass pointers to every field in exact SELECT order
		err := rows.Scan(
			&translation.Lexentry,
			&translation.Sense_Num,
			&translation.Sense,
			&translation.Written_Rep,
			&translation.TransList,
			&translation.Score,
			&translation.Is_Good,
			&translation.Imporatance,
		)
		translation.TranslationText()
		if err != nil {
			return Translation{}, fmt.Errorf("scan error: %w", err)
		}

	}

	if err := rows.Err(); err != nil {
		return Translation{}, err
	}

	return translation, nil
}

// Access helper when reading the result: avoids null sql strings
func (t Translation) TranslationText() {
	if t.Lexentry.Valid {
	} else {
		t.Lexentry.String = ""
	}
	if t.Sense_Num.Valid {
	} else {
		t.Sense_Num.String = ""
	}
	if t.Sense.Valid {
	} else {
		t.Sense.String = ""
	}
}
