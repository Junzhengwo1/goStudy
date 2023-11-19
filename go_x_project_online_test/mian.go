package main

import (
	"goStudy/go_x_project_online_test/config"
	"goStudy/go_x_project_online_test/router"
)

func main() {
	// 显示的初始化函数
	config.Init()
	err := router.Router().Run(config.Config.Port)
	if err != nil {
		return
	}

}
