package setting

import (
	"github.com/moomman/tour2-blog/internal/global"
	"github.com/moomman/tour2-blog/pkg/setting"
	"log"
	"sync"
)

//首先加载所有的配置文件

var (
	configPath = global.RootDir + "/config/app" //配置文件路径
	configName = "app"
	configType = "yml"
)

var once sync.Once

//会在配置文件读取之前加载
func init() {
	once.Do(func() {
		nSet := setting.NewSetting(configName, configType, configPath)
		if err := nSet.BindAll(&global.Setting); err != nil {
			log.Println(err)
			panic("配置文件绑定错误")
		}
	})
}
