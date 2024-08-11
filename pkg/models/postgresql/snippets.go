package postgresql

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"richardobaze.com/snippetbox/pkg/models"
)

// Define a snippetModel type which
type SnippetModel struct {
	DB *pgxpool.Pool
}

// This will insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
