package download

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite" // Pure Go SQLite driver (no CGO required)
)

const wikdictBaseURL = "https://download.wikdict.com/dictionaries/sqlite/2_2026-06/"

func GetDictionary(fromLang, toLang string) (*sql.DB, error) {
	pair := fmt.Sprintf("%s-%s", fromLang, toLang)
	filename := pair + ".sqlite3"

	// Store locally app directory
	localDir := "data/dictionaries"
	localPath := filepath.Join(localDir, filename)

	// Step 1: Check if already downloaded
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		log.Printf("Dictionary %s not found locally. Downloading from WikDict...\n", pair)
		downloadURL := fmt.Sprintf("%s%s.sqlite3", wikdictBaseURL, pair)

		if err := downloadAndDecompress(downloadURL, localPath, localDir); err != nil {
			return nil, fmt.Errorf("failed to fetch dictionary: %s %w", downloadURL, err)
		}
	}

	// Step 2: Open and query the local SQLite DB
	return sql.Open("sqlite", localPath)
}

func downloadAndDecompress(url, targetPath, dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned status: %s", resp.Status)
	}

	// Create local uncompressed target file
	out, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
