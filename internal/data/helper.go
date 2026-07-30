package data

import (
	"database/sql"
	"fmt"

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

// ReadColumnValues fetches all values from a specific column in a table.
func ReadColumnValues(db *sql.DB, tableName, columnName string, importance float32) ([]string, error) {
	// Construct the query. (Use parameter substitution for values, but table/column
	// names in SQL must be formatted into the query string).
	query := fmt.Sprintf("SELECT %s FROM %s WHERE importance >= %f LIMIT 50", columnName, tableName, importance)

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

func GetTranslations(db *sql.DB, importance float32) ([]Translation, error) {
	// Query all 7 columns from the translation table
	query := fmt.Sprintf("SELECT * FROM translation WHERE importance >= %f LIMIT 200", importance)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var translations []Translation

	for rows.Next() {
		var t Translation

		// Pass pointers to every field in exact SELECT order
		err := rows.Scan(
			&t.Lexentry,
			&t.Sense_Num,
			&t.Sense,
			&t.Written_Rep,
			&t.TransList,
			&t.Score,
			&t.Is_Good,
			&t.Imporatance,
		)
		t.TranslationText()
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		translations = append(translations, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return translations, nil
}

// Access helper when reading the result:
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
