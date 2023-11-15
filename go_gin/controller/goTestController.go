package controller

import (
	"github.com/gin-gonic/gin"
	"goStudy/go_gorm/dao"
	"strconv"
)

type GoTestController struct {
}

func (GoTestController) FindOne(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	query := dao.Query(id)
	Success(c, 200, "success", query)
}
