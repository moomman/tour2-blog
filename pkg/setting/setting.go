package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	viper *viper.Viper
}

//将configPath设置为可变参数，以便从多处查找配置文件

func NewSetting(configName, configType string, configPath ...string) *Setting {
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	for i := range configPath {
		if configPath[i] != "" {
			v.AddConfigPath(configPath[i])
		}
	}
	err := v.ReadInConfig()
	if err != nil {
		panic("找不到配置文件")
	}
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件改变： ", in.Name)
	})
	//监视配置文件的变化
	v.WatchConfig()

	return &Setting{v}
}

//存储配置记录
var all interface{}

func (s *Setting) BindAll(v interface{}) error {
	err := s.viper.Unmarshal(v)
	if err != nil {
		return err
	}
	all = v
	return nil
}
