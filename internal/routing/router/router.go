package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moomman/tour2-blog/internal/routing"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	//从根目录向下扩展
	root := engine.Group("")
	{
		routing.Group.Article.Init(root)
	}
	return engine
}
