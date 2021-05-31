package models

import (
	"database/sql"
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(db *sql.DB, content string) (err error) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			log.Println(err)
			tx.Rollback()
		}
	}()

	cmd := `insert into todos (
		content,
		user_id,
		created_at) values ($1, $2, $3)`
	_, err = tx.Exec(cmd, content, u.ID, time.Now())

	return err
}

func GetTodo(db *sql.DB, id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos
	where id = $1`
	todo = Todo{}

	err = db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

func GetTodos(db *sql.DB) (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser(db *sql.DB) (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos
	where user_id = $1`

	rows, err := db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			log.Println(err)
			tx.Rollback()
		}
	}()

	cmd := `update todos set content = $1, user_id = $2
	where id = $3`
	_, err = tx.Exec(cmd, t.Content, t.UserID, t.ID)

	return err
}

func (t *Todo) DeleteTodo(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			log.Println(err)
			tx.Rollback()
		}
	}()

	cmd := `delete from todos where id = $1`
	_, err = tx.Exec(cmd, t.ID)

	return err
}
