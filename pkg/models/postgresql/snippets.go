package postgresql

import (
	"context"
	"errors"
	"fmt"

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
		SELECT id, title, content, created, expires FROM public.snippets
		WHERE expires > NOW() AND id = @id
		`

	snip := &models.Snippet{}

	err := m.DB.QueryRow(context.Background(), insertStmt, args).Scan(
		&snip.ID,
		&snip.Title,
		&snip.Content,
		&snip.Created,
		&snip.Expires)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function check for that
		// error specifically, and return our own models.ErrNoRecord error
		// instead.
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return snip, nil
}

// This will return up to 10 of the most recently created snippets
func (m *SnippetModel) Latest() ([]models.Snippet, error) {
	insertStmt := `
		SELECT id, title, content, created, expires FROM public.snippets
		WHERE expires > NOW() ORDER BY created DESC LIMIT 10;
		`

	rows, err := m.DB.Query(context.Background(), insertStmt)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	defer rows.Close()

	snips, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Snippet])
	if err != nil {
		return nil, fmt.Errorf("unable to Collect Rows: %w", err)
	}

	return snips, nil
}
