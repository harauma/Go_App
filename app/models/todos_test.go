package models

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	u := User{ID: 1}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`insert into todos (
		content,
		user_id,
		created_at) values ($1, $2, $3)`)).
		WithArgs("testTodo", u.ID, AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// モック化されたDBを用いてテスト対象関数を実行
	if err = u.CreateTodo(db, "testTodo"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestGetTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()
	time := time.Now()
	todo := Todo{1, "test", 1, time}
	columns := []string{"id", "content", "user_id", "created_at"}
	mock.ExpectQuery(regexp.QuoteMeta(`select id, content, user_id, created_at from todos where id = $1`)).
		WithArgs(1).WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "test", 1, time))

	resultTodo, err := GetTodo(db, 1)
	if err != nil {
		t.Fatalf("failed to get todo: %s", err)
	}

	if todo != resultTodo {
		t.Fatalf("It's a different value than the expected value.: %v", resultTodo)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWerMet(): %s", err)
	}
}

func TestGetTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()
	time := time.Now()
	todos := []Todo{
		{ID: 1, Content: "test", UserID: 1, CreatedAt: time},
		{ID: 2, Content: "test2", UserID: 1, CreatedAt: time},
	}
	columns := []string{"id", "content", "user_id", "created_at"}
	mock.ExpectQuery(regexp.QuoteMeta("select id, content, user_id, created_at from todos")).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "test", 1, time).AddRow(2, "test2", 1, time))

	resultTodos, err := GetTodos(db)
	if err != nil {
		t.Fatalf("failed to get todos: %s", err)
	}

	if diff := cmp.Diff(todos, resultTodos); diff != "" {
		t.Errorf("User value is mismatch (-tom +tom2):\n%s", diff)
	}

	// mock定義の期待操作が順序道理に実行されたか検査
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWerMet(): %s", err)
	}

}

func TestGetTodosByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()
	time := time.Now()
	u := User{ID: 1}
	todos := []Todo{
		{ID: 1, Content: "test", UserID: 1, CreatedAt: time},
		{ID: 2, Content: "test2", UserID: 1, CreatedAt: time},
	}
	columns := []string{"id", "content", "user_id", "created_at"}
	mock.ExpectQuery(regexp.QuoteMeta(`select id, content, user_id, created_at from todos where user_id = $1`)).
		WithArgs(1).WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "test", 1, time).AddRow(2, "test2", 1, time))

	resultTodos, err := u.GetTodosByUser(db)
	if err != nil {
		t.Fatalf("failed to get todo: %s", err)
	}

	if diff := cmp.Diff(todos, resultTodos); diff != "" {
		t.Errorf("User value is mismatch (-tom +tom2):\n%s", diff)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWerMet(): %s", err)
	}
}
func TestUpdateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	time := time.Now()
	todo := Todo{ID: 1, Content: "testTodo", UserID: 1, CreatedAt: time}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`update todos set content = $1, user_id = $2
	where id = $3`)).
		WithArgs("testTodo", todo.UserID, todo.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// モック化されたDBを用いてテスト対象関数を実行
	if err = todo.UpdateTodo(db); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestDeleteTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	time := time.Now()
	todo := Todo{ID: 1, Content: "testTodo", UserID: 1, CreatedAt: time}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`delete from todos where id = $1`)).
		WithArgs(todo.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// モック化されたDBを用いてテスト対象関数を実行
	if err = todo.DeleteTodo(db); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}
