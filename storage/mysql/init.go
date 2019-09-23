package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"web-layout/utils/mysql"
)

type MySQL struct {
	config mysql.Config
	conn   *sql.DB
}

func NewMySQL(c mysql.Config) (*MySQL, error) {
	conn, err := c.Connect()
	if err != nil {
		return nil, err
	}

	return &MySQL{
		config: c,
		conn:   conn,
	}, nil
}
