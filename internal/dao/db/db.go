package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	db "github.com/moomman/tour2-blog/internal/dao/db/sqlc"
)

type DB struct {
	db.Store
}

func Init(dataSourceSet string) *DB {
	conn, err := sql.Open("mysql", dataSourceSet)
	if err != nil {
		panic(err)
	}
	return &DB{
		Store: db.SqlStore{
			Queries: db.New(conn),
			DB:      conn,
		},
	}
}
