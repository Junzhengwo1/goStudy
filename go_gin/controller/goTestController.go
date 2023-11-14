package controller

import (
	"github.com/gin-gonic/gin"
	"goStudy/go_gorm/gormdao"
	"strconv"
)

type GoTestController struct {
}

func (g GoTestController) FindOne(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	query := gormdao.Query(id)
	Success(c, 200, "success", query)
}
