package reply

import (

	db "github.com/moomman/tour2-blog/internal/dao/db/sqlc"
)

type ListArticle struct {
	[]*db.Article `json:"db_article"`
}
