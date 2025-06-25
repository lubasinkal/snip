package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lubasinkal/snip/internal/models"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func init() {
	var err error

	// Get database path
	dbPath := getDBPath()

	// Ensure directory exists
	dbDir := filepath.Dir(dbPath)
	err = os.MkdirAll(dbDir, 0755)
	if err != nil {
		panic(fmt.Sprintf("Failed to create database directory: %v", err))
	}

	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS snippets (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        tags TEXT,
        content TEXT,
        created_at DATETIME
    )`)
	if err != nil {
		panic(err)
	}
}

func SaveSnippet(s models.Snippet) (int64, error) {
	res, err := db.Exec(`INSERT INTO snippets (title, tags, content, created_at) VALUES (?, ?, ?, ?)`,
		s.Title, strings.Join(s.Tags, ","), s.Content, s.CreatedAt.Format("2006-01-02 15:04:05"))
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// ListAllSnippets returns all snippets from the database
func ListAllSnippets() ([]models.Snippet, error) {
	rows, err := db.Query(`SELECT id, title, tags, content, created_at FROM snippets ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []models.Snippet
	for rows.Next() {
		var s models.Snippet
		var tagsStr string
		var createdAtStr string

		err := rows.Scan(&s.ID, &s.Title, &tagsStr, &s.Content, &createdAtStr)
		if err != nil {
			return nil, err
		}

		// Parse tags
		if tagsStr != "" {
			s.Tags = strings.Split(tagsStr, ",")
		}

		// Parse created_at
		if createdAtStr != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr); err == nil {
				s.CreatedAt = parsedTime
			} else {
				// Try alternative format
				if parsedTime, err := time.Parse(time.RFC3339, createdAtStr); err == nil {
					s.CreatedAt = parsedTime
				} else {
					s.CreatedAt = time.Now() // fallback
				}
			}
		}

		snippets = append(snippets, s)
	}

	return snippets, nil
}

// GetSnippetByID returns a single snippet by its ID
func GetSnippetByID(id int) (*models.Snippet, error) {
	var s models.Snippet
	var tagsStr string
	var createdAtStr string

	err := db.QueryRow(`SELECT id, title, tags, content, created_at FROM snippets WHERE id = ?`, id).
		Scan(&s.ID, &s.Title, &tagsStr, &s.Content, &createdAtStr)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("snippet with ID %d not found", id)
		}
		return nil, err
	}

	// Parse tags
	if tagsStr != "" {
		s.Tags = strings.Split(tagsStr, ",")
	}

	// Parse created_at
	if createdAtStr != "" {
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr); err == nil {
			s.CreatedAt = parsedTime
		} else {
			// Try alternative format
			if parsedTime, err := time.Parse(time.RFC3339, createdAtStr); err == nil {
				s.CreatedAt = parsedTime
			} else {
				s.CreatedAt = time.Now() // fallback
			}
		}
	}

	return &s, nil
}

// SearchSnippets searches for snippets by query in title, tags, or content
func SearchSnippets(query string, tagFilter string) ([]models.Snippet, error) {
	var rows *sql.Rows
	var err error

	query = strings.ToLower(query)

	if tagFilter != "" {
		// Search with tag filter - use word boundaries to match exact tags
		rows, err = db.Query(`
			SELECT id, title, tags, content, created_at
			FROM snippets
			WHERE (LOWER(title) LIKE ? OR LOWER(content) LIKE ?)
			AND (LOWER(tags) LIKE ? OR LOWER(tags) LIKE ? OR LOWER(tags) LIKE ? OR LOWER(tags) = ?)
			ORDER BY created_at DESC`,
			"%"+query+"%", "%"+query+"%",
			strings.ToLower(tagFilter)+",%", "%,"+strings.ToLower(tagFilter)+",%", "%,"+strings.ToLower(tagFilter), strings.ToLower(tagFilter))
	} else {
		// Search without tag filter
		rows, err = db.Query(`
			SELECT id, title, tags, content, created_at
			FROM snippets
			WHERE LOWER(title) LIKE ? OR LOWER(tags) LIKE ? OR LOWER(content) LIKE ?
			ORDER BY created_at DESC`,
			"%"+query+"%", "%"+query+"%", "%"+query+"%")
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []models.Snippet
	for rows.Next() {
		var s models.Snippet
		var tagsStr string
		var createdAtStr string

		err := rows.Scan(&s.ID, &s.Title, &tagsStr, &s.Content, &createdAtStr)
		if err != nil {
			return nil, err
		}

		// Parse tags
		if tagsStr != "" {
			s.Tags = strings.Split(tagsStr, ",")
		}

		// Parse created_at
		if createdAtStr != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr); err == nil {
				s.CreatedAt = parsedTime
			} else {
				// Try alternative format
				if parsedTime, err := time.Parse(time.RFC3339, createdAtStr); err == nil {
					s.CreatedAt = parsedTime
				} else {
					s.CreatedAt = time.Now() // fallback
				}
			}
		}

		snippets = append(snippets, s)
	}

	return snippets, nil
}

// UpdateSnippet updates an existing snippet
func UpdateSnippet(s models.Snippet) error {
	_, err := db.Exec(`UPDATE snippets SET title = ?, tags = ?, content = ? WHERE id = ?`,
		s.Title, strings.Join(s.Tags, ","), s.Content, s.ID)
	return err
}

// DeleteSnippet removes a snippet by ID
func DeleteSnippet(id int) error {
	result, err := db.Exec(`DELETE FROM snippets WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("snippet with ID %d not found", id)
	}

	return nil
}

// getDBPath returns the path to the database file
func getDBPath() string {
	// Check for custom path in environment variable
	if dbPath := os.Getenv("SNIP_DB_PATH"); dbPath != "" {
		return dbPath
	}

	// Default to ~/.snipdb/snippets.db
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Fallback to temp directory
		return filepath.Join(os.TempDir(), "snip.db")
	}

	return filepath.Join(homeDir, ".snipdb", "snippets.db")
}
