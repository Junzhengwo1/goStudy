package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goStudy/go_xproject_chat/common"
	"goStudy/go_xproject_chat/dao"
	"strconv"
)

type UserController struct {
}

func (*UserController) Query(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	// service - > dao
	userDao := dao.UserDao{}
	fmt.Printf("type is %T \n", userDao)
	query := userDao.Query(id)
	common.Success(c, 200, "success", query)
}
