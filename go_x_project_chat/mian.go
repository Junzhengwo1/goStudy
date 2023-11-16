package main

import (
	"goStudy/go_x_project_chat/router"
	"goStudy/go_x_project_chat/util"
)

func main() {
	//util.InitConfig()

	err := router.Router().Run(util.ServerConf["port"].(string))
	if err != nil {
		return
	}

}
