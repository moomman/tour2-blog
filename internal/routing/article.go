package routing

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/moomman/tour2-blog/internal/api/v1"
)

type article struct {
}

func (article) Init(group *gin.RouterGroup) {
	routerGroup := group.Group("article")
	{
		routerGroup.GET("list", v1.Group.Article.ListArticle)
	}
}
