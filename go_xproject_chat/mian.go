package main

import "goStudy/go_xproject_chat/router"

const PORT = ":8899"

func main() {

	err := router.Router().Run(PORT)
	if err != nil {
		return
	}

}
