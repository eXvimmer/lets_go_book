package models

import (
	"database/sql"
	"errors"
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

// return a specifc snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	/* NOTE:
	 * The placeholder parameter syntax differs depending on your database.
	 * MySQL, SQL Server and SQLite use the ? notation, but PostgreSQL uses the
	 * $N notation, e.g. $1, $2 ...
	 */
	stmt := `
    SELECT id, title, content, created, expires
    FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?
  `

	row := m.DB.QueryRow(stmt, id)

	s := &Snippet{}
	/* NOTE:
	 * Notice that the arguments to row.Scan are *pointers* to the place you want
	 * to copy the data into, and the number of arguments must be exactly the
	 * same as the number of columns returned by your statement.
	 */
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// TODO: return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
