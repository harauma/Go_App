package models

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

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
	todos := []Todo{{ID: 1, Content: "test", UserID: 1, CreatedAt: time}, {ID: 2, Content: "test2", UserID: 1, CreatedAt: time}}
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
	todos := []Todo{{ID: 1, Content: "test", UserID: 1, CreatedAt: time}, {ID: 2, Content: "test2", UserID: 1, CreatedAt: time}}
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
