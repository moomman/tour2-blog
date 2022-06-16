package setting

import (
	"github.com/moomman/tour2-blog/internal/dao"
	"github.com/moomman/tour2-blog/internal/dao/db"
	"github.com/moomman/tour2-blog/internal/global"
)

type mDao struct {
}

func (mDao) Init() {
	dao.Group.DB = db.Init(global.Setting.Mysql.DataSourceName)
}
