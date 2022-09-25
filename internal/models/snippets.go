package models

import (
	"database/sql"
	"time"
)

// Snippet type holds the data for an individual snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

//  Insert a new snippet to the database and return it's id and an error
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	stmt := `
    INSERT INTO snippets (title, content, created, expires)
    VALUES (
      ?,
      ?,
      UTC_TIMESTAMP(),
      DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY)
    );
  `
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil // NOTE: I'm not sure about returning int type for id
}

// TODO: return a specifc snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// TODO: return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
