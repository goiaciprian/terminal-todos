-- name: GetTodos :many
select * from todos;

-- name: CountTodos :one
select count(*) from todos;

-- name: GetTodosById :one
select * from todos where id = ?;

-- name: CreateTodo :one
insert into todos (title, description)
values (?, ?)
returning *;

-- name: CompleteTodo :one
update todos
set completed = 1,
completed_at = datetime('now', 'localtime')
where id = ?
returning *;

-- name: UncompleteTodo :one
update todos
set completed = 0,
completed_at = null
where id = ?
returning *;

-- name: UpdateTodo :one
update todos
set title = ?,
description = ?
where id = ?
returning *;

-- name: DeleteTodo :one
delete from todos
where id = ?
returning *;