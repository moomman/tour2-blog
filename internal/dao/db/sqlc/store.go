package db

import "database/sql"

//提高可扩展性

type Store interface {
	Querier
	TXer
}

//注册所有的事务

type TXer interface {
}

type SqlStore struct {
	*Queries
	DB *sql.DB
}
