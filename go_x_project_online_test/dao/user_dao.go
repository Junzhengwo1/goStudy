package dao

import (
	"fmt"
	"goStudy/go_x_project_online_test/model/pojo"
	"goStudy/go_x_project_online_test/util"
	"log"
)

type UserDao struct {
}

func (*UserDao) Query(id int) *pojo.UserBasic {
	// 查询单条数据
	var user *pojo.UserBasic
	tx := util.Db.Where("id = ?", id).First(&user)
	err := tx.Error
	if err != nil {
		log.Fatalln(err)
	}
	affected := tx.RowsAffected
	fmt.Println("affected=", affected)
	return user
}
