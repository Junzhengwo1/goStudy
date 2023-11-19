package main

import (
	"goStudy/go_x_project_online_test/config"
	"goStudy/go_x_project_online_test/router"
)

func main() {

	err := router.Router().Run(config.Config.Port)
	if err != nil {
		return
	}

}
