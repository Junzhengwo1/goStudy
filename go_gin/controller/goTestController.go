package handler

import (
	"github.com/gin-gonic/gin"
	"goStudy/go_gin/util/self_logger"
	"goStudy/go_gorm/dao"
	"strconv"
)

type GoTestHandler struct {
}

func (GoTestHandler) FindOne(c *gin.Context) {
	idStr := c.Query("id")
	self_logger.WriteLog("idStr", "user")
	id, _ := strconv.Atoi(idStr)
	query := dao.Query(id)
	Success(c, 200, "success", query)
}
