package main

import "goStudy/go_gin/router"

func main() {

	userRouter := router.UserRouter()

	err := userRouter.Run(":8089")
	if err != nil {
		return
	}
}
