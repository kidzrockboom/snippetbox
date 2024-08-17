package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"richardobaze.com/snippetbox/pkg/models"
)

// Define a snippetModel type which
type SnippetModel struct {
	DB *pgxpool.Pool
}

// This will insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (pgconn.CommandTag, error) {
	args := pgx.NamedArgs{
		"title":   title,
		"content": content,
		"expires": expires,
	}

	insertStmt := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES(
		@title, 
		@content,
		TIMEZONE('UTC', NOW()), 
		DATE_ADD(TIMEZONE('UTC', NOW()), INTERVAL '@expires DAYS')
		)`

	result, err := m.DB.Exec(context.Background(), insertStmt, args)
	if err != nil {
		return pgconn.NewCommandTag("empty"), err
	}

	return result, nil
}

// This will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
