// todos_dao_pg.go
package main

import (
    "database/sql"
    _ "github.com/jackc/pgx/v5/stdlib"
)

type TodoDaoPGImpl struct {
    conn *sql.DB
}

func NewPGTodoDao(conn *sql.DB) TodoDao {
    return &TodoDaoPGImpl{conn: conn}
}

func (dao *TodoDaoPGImpl) GetAll() ([]*Todo, error) {
    rows, err := dao.conn.Query(`
        SELECT t.id, t.title, t.completed, t.created_at, t.updated_at, t.deleted_at
        FROM todos t
        WHERE t.deleted_at IS NULL 
        ORDER BY t.created_at desc
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    todos := []*Todo{}
    for rows.Next() {
        todo := &Todo{}
        err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
        if err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}
func (dao *TodoDaoPGImpl) Get(id string) (*Todo, error) {
	todo := &Todo{}
	err := dao.conn.QueryRow(`
	 SELECT 
	  t.id
	  , t.title
	  , t.completed
	  , t.created_at
	  , t.updated_at
	 FROM todos t
	 WHERE t.id = $1
	`, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
	 return nil, err
	}
	return todo, nil
   }

func (dao *TodoDaoPGImpl) GetHistory() ([]*Todo, error) {
    rows, err := dao.conn.Query(`
        SELECT t.id, t.title, t.completed, t.created_at, t.updated_at, t.deleted_at
        FROM todos t
        ORDER BY t.deleted_at DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    todos := []*Todo{}
    for rows.Next() {
        todo := &Todo{}
        err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
        if err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

func (dao *TodoDaoPGImpl) Create(todo *Todo) error {
    _, err := dao.conn.Exec("INSERT INTO todos (id, title, completed, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())",
        todo.ID, todo.Title, todo.Completed)
    return err
}

func (dao *TodoDaoPGImpl) Update(todo *Todo) error {
    _, err := dao.conn.Exec("UPDATE todos SET title = $1, completed = $2, updated_at = now() WHERE id = $3",
        todo.Title, todo.Completed, todo.ID)
    return err
}

func (dao *TodoDaoPGImpl) Delete(id string) error {
    _, err := dao.conn.Exec("UPDATE todos SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1", id)
    return err
}

func (dao *TodoDaoPGImpl) Done(id string) error {
	_, err := dao.conn.Exec("UPDATE todos SET completed = NOT completed, updated_at = now() WHERE id = $1", id)
	return err
   }