package dao

import (
	"goStudy/go_x_project_chat/model/pojo"
	"goStudy/go_x_project_chat/util"
	"log"
)

type UserDao struct {
}

func (*UserDao) Query(id int) pojo.User {
	// 查询单条数据
	var user pojo.User
	err := util.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Fatalln(err)
	}
	return user
}
