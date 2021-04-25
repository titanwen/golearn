package chapter02

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pkgerrors "github.com/pkg/errors"
)

func TestSqlClient_QueryRaw(t *testing.T) {
	t.Run("query with raw error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("sqlmock new error: %v", err)
		}

		// mock sql select
		mock.ExpectQuery("select").WillReturnError(sql.ErrNoRows)
		mock.ExpectClose()

		sqlClient := NewSqlClientFromDB(db)
		defer func() {
			if err := sqlClient.Close(); err != nil {
				t.Errorf("close sql fail: %s", err.Error())
			}
		}()

		rows, err := sqlClient.QueryRaw("select * from task")
		if err != nil {
			// 判断是否为sql.ErrNoRows => false，无法判断，只能通过关键字匹配
			t.Logf("query sql fail, is error no rows: %v", errors.Is(err, sql.ErrNoRows))

			// 打印错误 => 原始字符串
			t.Logf("err: %+v", err)
			return
		}
		defer func() {
			if err := rows.Close(); err != nil {
				t.Fatalf("close sql rows fail: %v", err)
			}
		}()
	})
}

func TestSqlClient_QueryWithWrap(t *testing.T) {
	t.Run("query with wrap error", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("sqlmock new error: %v", err)
		}

		// mock sql select
		mock.ExpectQuery("select").WillReturnError(sql.ErrNoRows)
		mock.ExpectClose()

		sqlClient := NewSqlClientFromDB(db)
		defer func() {
			if err := sqlClient.Close(); err != nil {
				t.Errorf("close sql fail: %s", err.Error())
			}
		}()

		rows, err := sqlClient.QueryWithWrap("select * from task")
		if err != nil {
			// 判断是否为sql.ErrNoRows => true 可以判断
			t.Logf("query sql fail, is error no rows: %v", errors.Is(err, sql.ErrNoRows))

			// 打印错误 => 原始字符串
			t.Logf("err: %+v", err)
			return
		}
		defer func() {
			if err := rows.Close(); err != nil {
				t.Fatalf("close sql rows fail: %v", err)
			}
		}()
	})
}

func TestSqlClient_QueryWithPkgWrap(t *testing.T) {
	t.Run("query with pkg wrap error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("sqlmock new error: %v", err)
		}

		// mock sql select
		mock.ExpectQuery("select").WillReturnError(sql.ErrNoRows)
		mock.ExpectClose()

		sqlClient := NewSqlClientFromDB(db)
		defer func() {
			if err := sqlClient.Close(); err != nil {
				t.Errorf("close sql fail: %s", err.Error())
			}
		}()

		rows, err := sqlClient.QueryWithPkgWrap("select * from task")
		if err != nil {
			// 判断是否为sql.ErrNoRows => true 可以判断
			t.Logf("query sql fail, is error no rows: %v", pkgerrors.Is(err, sql.ErrNoRows))

			// 打印错误 => 堆栈信息
			t.Logf("err: %+v", err)
			return
		}
		defer func() {
			if err := rows.Close(); err != nil {
				t.Fatalf("close sql rows fail: %v", err)
			}
		}()
	})
}
