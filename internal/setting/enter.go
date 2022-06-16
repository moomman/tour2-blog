package setting

type group struct {
	MDao mDao
}

var Group = new(group)

//初始化所有的组件

func AllInit() {
	Group.MDao.Init()
}
