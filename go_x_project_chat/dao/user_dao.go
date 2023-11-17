package dao

import (
	"fmt"
	"goStudy/go_x_project_chat/model/pojo"
	"goStudy/go_x_project_chat/util"
	"log"
)

type UserDao struct {
}

func (*UserDao) Query(id int) *pojo.User {
	// 查询单条数据
	var user *pojo.User
	tx := util.Db.Where("id = ?", id).First(&user)
	err := tx.Error
	if err != nil {
		log.Fatalln(err)
	}
	affected := tx.RowsAffected
	fmt.Println(affected)
	return user
}
