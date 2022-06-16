package dao

import "github.com/moomman/tour2-blog/internal/dao/db"

type group struct {
	DB *db.DB
}

var Group = new(group)
