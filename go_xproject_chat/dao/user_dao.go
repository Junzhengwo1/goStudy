package dao

import (
	"fmt"
	"goStudy/go_xproject_chat/model/pojo"
	"goStudy/go_xproject_chat/util"
)

type UserDao struct {
}

func (*UserDao) Query(id int) pojo.User {
	// 查询单条数据
	var user pojo.User
	err := util.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user is %v", user)
	return user
}
