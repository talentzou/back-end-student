package core

import (
	"back-end/core/internal"
	"back-end/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
// // 读取配置文件路由
// func ViperRouter() *viper.Viper {
// 	var config string = "route.yaml"
// 	v := viper.New()
// 	v.SetConfigFile(config)
// 	v.SetConfigType("yaml")
// 	err := v.ReadInConfig() // Find and read the config file
// 	if err != nil {         // Handle errors reading the config file
// 		panic(fmt.Errorf("fatal error config file: %w", err))
// 	}
// 	// 监听配置文件的改变
// 	viper.OnConfigChange(func(e fsnotify.Event) {
// 		fmt.Println("Config file changed:", e.Name)
// 		fmt.Println("config file changed:", e.Name)
// 		if err = v.Unmarshal(&global.Global_Web_Route); err != nil {
// 			fmt.Println(err)
// 		}
// 	})
// 	if err = v.Unmarshal(&global.Global_Web_Route); err != nil {
// 		panic(err)
// 	}
// 	viper.WatchConfig()
// 	return v
// }

// 读取路由
func ViperServer() *viper.Viper {
	var config string = internal.ConfigDefaultFile
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.WatchConfig()
	// 监听配置文件的改变
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Global_Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Global_Config); err != nil {
		panic(err)
	}

	return v
}
