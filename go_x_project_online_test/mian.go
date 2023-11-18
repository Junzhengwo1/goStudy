package main

import (
	"goStudy/go_x_project_online_test/router"
	"goStudy/go_x_project_online_test/util"
)

func main() {
	//util.InitConfig()

	err := router.Router().Run(util.ServerConf["port"].(string))
	if err != nil {
		return
	}

}
