package conf

import (
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"soulapi/global"
)

func init() {
	configPath := beego.AppConfig.String("config_path")
	if configPath == "" {
		panic("未设置配置文件路径")
	}
	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GLOBAL_CONFIG); err != nil {
			panic(err)
		}
	})

	if err := v.Unmarshal(&global.GLOBAL_CONFIG); err != nil {
		panic(err)
	}
}
