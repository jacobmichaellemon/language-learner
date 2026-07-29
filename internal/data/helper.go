package data

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// ReadColumnValues fetches all values from a specific column in a table.
func ReadColumnValues(db *sql.DB, tableName, columnName string) ([]string, error) {
	// Construct the query. (Use parameter substitution for values, but table/column
	// names in SQL must be formatted into the query string).
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT 50", columnName, tableName)

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
