package dao

import (
	"fmt"
	"goStudy/go_xproject_chat/config"
	"goStudy/go_xproject_chat/model/pojo"
)

type UserDao struct {
}

func (*UserDao) Query(id int) pojo.User {
	// 查询单条数据
	var user pojo.User
	err := config.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	return user
}
