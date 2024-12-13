package todo

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"strconv"
	"terminal-todos/internal/config"
	"terminal-todos/internal/database"
	"time"

	"github.com/aquasecurity/table"
)

var queries database.Queries
var db *sql.DB

func init() {
	_db, err := database.Open(config.Instance.DATABASE_FILE_PATH)
	if err != nil {
		return
	}

	db = _db
	queries = *database.New(db)
}

func Count() (int64, error) {
	ctx := context.Background()
	return queries.CountTodos(ctx)
}

func CloseDb() {
	defer db.Close()
}

func GetAll(qType, order string) ([]database.Todo, error) {
	defer db.Close()
	ctx := context.Background()
	return queries.GetTodos(ctx, qType, order)
}

func Create(title, description string) ([]database.Todo, error) {
	defer db.Close()
	ctx := context.Background()

	_todo, err := queries.CreateTodo(ctx, database.CreateTodoParams{
		Title: title,
		Description: sql.NullString{
			String: description,
			Valid:  true,
		},
	})
	if err != nil {
		return []database.Todo{}, err
	}
	return []database.Todo{_todo}, nil
}

func Complete(id int64) ([]database.Todo, error) {
	defer db.Close()
	ctx := context.Background()

	_, err := queries.GetTodosById(ctx, id)
	if err != nil {
		return []database.Todo{}, errors.New("todo not found")
	}

	todo, err := queries.CompleteTodo(ctx, id)
	return []database.Todo{todo}, err
}

func Uncomplete(id int64) ([]database.Todo, error) {
	defer db.Close()
	ctx := context.Background()

	_, err := queries.GetTodosById(ctx, id)
	if err != nil {
		return []database.Todo{}, errors.New("todo not found")
	}

	todo, err := queries.UncompleteTodo(ctx, id)
	return []database.Todo{todo}, err
}

func Update(id int64, title, description string) ([]database.Todo, error) {
	defer db.Close()
	ctx := context.Background()

	existing, err := queries.GetTodosById(ctx, id)
	if err != nil {
		return []database.Todo{}, err
	}

	newTitle := existing.Title
	newDesc := existing.Description.String

	if len(title) > 0 {
		newTitle = title
	}

	if len(description) > 0 {
		newDesc = description
	}

	todo, err := queries.UpdateTodo(ctx, database.UpdateTodoParams{
		Title: newTitle,
		Description: sql.NullString{
			String: newDesc,
			Valid:  true,
		},
		ID: id,
	})
	if err != nil {
		return []database.Todo{}, err
	}

	return []database.Todo{todo}, nil
}

func Delete(id int64) ([]database.Todo, error) {
	defer db.Close()
	ctx := context.Background()

	_, err := queries.GetTodosById(ctx, id)
	if err != nil {
		return []database.Todo{}, err
	}

	todo, err := queries.DeleteTodo(ctx, id)
	if err != nil {
		return []database.Todo{}, err
	}
	return []database.Todo{todo}, nil
}

func Display(todos *[]database.Todo) {
	table := table.New(os.Stdout)
	table.SetHeaders("#", "Title", "Description", "Created At", "Completed", "Completed At")
	for _, t := range *todos {
		id := strconv.Itoa(int(t.ID))
		completed := ""
		completedAt := ""
		if t.Completed.Int64 != 0 {
			completed = "true"
			completedAt = t.CompletedAt.Time.Format(time.RFC3339)
		}
		table.AddRow(id, t.Title, t.Description.String, t.CreatedAt.Format(time.RFC3339), completed, completedAt)
	}

	table.Render()
}
