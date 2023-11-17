package controller

import (
	"github.com/gin-gonic/gin"
	"goStudy/go_x_project_chat/common"
	"goStudy/go_x_project_chat/service"
	"strconv"
)

type UserController struct {
}

func (*UserController) Query(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	// service - > dao
	userService := service.UserService{}
	query := userService.Query(id)
	common.Success(c, query)
}
