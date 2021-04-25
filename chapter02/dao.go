// Package chapter02
// 问题：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 答：按照毛老师在视频中讲的内容，我觉得应该分两种情况
// 1. 若此dao module是属于基础库，供模块或者多人使用，则应该直接返回原始错误信息，即 QueryRaw
// 2. 若此dao为application的一部分，即dao只属于这个application，则应该返回wrap信息，可以用 QueryWithWrap 或者 QueryWithPkgWrap
//
// 运行以下命令，查看三种错误信息的输出
// go test -v -failfast=false ./...
//
package chapter02

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

// SqlClient to manager mysql
type SqlClient struct {
	db *sql.DB
}

// NewSqlClientFromDB actually for mock
func NewSqlClientFromDB(db *sql.DB) *SqlClient {
	return &SqlClient{db: db}
}

// QueryRaw 通过添加信息，重新返回新的error
func (s *SqlClient) QueryRaw(querySql string) (*sql.Rows, error) {
	rows, err := s.db.Query(querySql)
	if err != nil {
		return nil, fmt.Errorf("query raw fail: %v", err)
	}
	return rows, nil
}

// QueryWithWrap errors包支持的wrap
func (s *SqlClient) QueryWithWrap(querySql string) (*sql.Rows, error) {
	rows, err := s.db.Query(querySql)
	if err != nil {
		return nil, fmt.Errorf("query with wrap fail: %w", err)
	}
	return rows, nil
}

// QueryWithWrap pkg/errors支持的wrap
func (s *SqlClient) QueryWithPkgWrap(querySql string) (*sql.Rows, error) {
	rows, err := s.db.Query(querySql)
	if err != nil {
		return nil, errors.Wrap(err, "query with pkg wrap fail")
	}
	return rows, nil
}

// Close close connection
func (s *SqlClient) Close() error {
	return s.db.Close()
}
