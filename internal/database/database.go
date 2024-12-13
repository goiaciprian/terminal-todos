package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		fmt.Printf("Error opening the db: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("No response from the db: %s\n", err)
	}

	return db, err
}

func getTodos(qType, order string) string {
	switch qType {
	case "completed":
		return fmt.Sprintf(`-- name: GetTodos :many
select id, title, description, created_at, completed, completed_at from todos where completed = 1 order by created_at %s
`, order)
	case "uncompleted":
		return fmt.Sprintf(`-- name: GetTodos :many
select id, title, description, created_at, completed, completed_at from todos where completed = 0 order by created_at %s
`, order)
	default:
		return fmt.Sprintf(`-- name: GetTodos :many
select id, title, description, created_at, completed, completed_at from todos order by created_at %s
`, order)
	}
}

func (q *Queries) GetTodos(ctx context.Context, qType, sort string) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getTodos(qType, sort))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.Completed,
			&i.CompletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
