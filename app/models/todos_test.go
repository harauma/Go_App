package models

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func init() {

}

func TestGetTodos(t *testing.T) {
	// 相対パスを再設定する
	// apath, _ := filepath.Abs("../")
	// os.Chdir(apath)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()
	columns := []string{"id", "content", "user_id", "created_at"}
	mock.ExpectQuery(regexp.QuoteMeta("select id, content, user_id, created_at from todos")).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "test", 1, time.Now()))

	// テスト対象関数call
	t.Log("gettodos")
	todos, err := GetTodos(db)
	if err != nil {
		t.Fatalf("failed to get todos: %s", err)
	}
	t.Logf("%v", todos)

	// mock定義の期待操作が順序道理に実行されたか検査
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWerMet(): %s", err)
	}

}
