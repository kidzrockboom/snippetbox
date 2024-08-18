package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"richardobaze.com/snippetbox/pkg/models"
)

// Define a snippetModel type which
type SnippetModel struct {
	DB *pgxpool.Pool
}

// This will insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// Variable to store the ID of the query
	var id int

	args := pgx.NamedArgs{
		"title":   title,
		"content": content,
		"expires": expires,
	}

	insertStmt := `
		INSERT INTO public.snippets (title, content, created, expires)
		VALUES(
		@title, 
		@content,
		TIMEZONE('UTC', NOW()), 
		DATE_ADD(TIMEZONE('UTC', NOW()), INTERVAL '1 DAYS' * @expires)
		) RETURNING id;`

	err := m.DB.QueryRow(context.Background(), insertStmt, args).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// This will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	args := pgx.NamedArgs{
		"id": id,
	}

	insertStmt := `
		SELECT id, title, content, created, expires, FROM public.snippets
		WHERE expires > NOW() AND id = @id
		);`

	snip := &models.Snippet{}

	err := m.DB.QueryRow(context.Background(), insertStmt, args).Scan()
	if err != nil {
		return snip, err
	}

	return nil, nil
}

// This will return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
