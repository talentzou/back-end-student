package core

import (
	"fmt"

	"back-end/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	var config string = "routeConfig.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// 监听配置文件的改变
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Global_Font_End_Route); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Global_Font_End_Route); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	return v
}
