package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"web-layout/utils/mysql"
)

type mysqlDB struct {
	config mysql.Config
	conn   *sql.DB
}

func MySQL(c mysql.Config) (*mysqlDB, error) {
	conn, err := c.Connect()
	if err != nil {
		return nil, err
	}

	return &mysqlDB{
		config: c,
		conn:   conn,
	}, nil
}
