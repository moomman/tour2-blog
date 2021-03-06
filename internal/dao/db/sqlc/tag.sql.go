// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: tag.sql

package db

import (
	"context"
)

const getAllTag = `-- name: GetAllTag :many
select id, article_id, tag_name from tag
`

func (q *Queries) GetAllTag(ctx context.Context) ([]*Tag, error) {
	rows, err := q.db.QueryContext(ctx, getAllTag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.ArticleID, &i.TagName); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
