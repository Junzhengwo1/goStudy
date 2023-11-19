package dao

import (
	"fmt"
	"goStudy/go_x_project_online_test/config"
	"goStudy/go_x_project_online_test/model/pojo"
	"log"
)

type UserDao struct {
}

func (*UserDao) Query(id int) *pojo.UserBasic {
	// 查询单条数据
	var user *pojo.UserBasic
	tx := config.Db.Where("id = ?", id).First(&user)
	err := tx.Error
	if err != nil {
		log.Fatalln(err)
	}
	affected := tx.RowsAffected
	fmt.Println("affected=", affected)
	return user
}
