package dao

import (
	"fmt"
	"goStudy/go_gorm/config"
	"goStudy/go_gorm/pojo"
)

func Query(id int) pojo.GoTest {
	// 查询单条数据
	var test pojo.GoTest
	err := config.Db.Where("id = ?", id).First(&test).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
	return test
}

// todo 其他操作

func AddUser(id int) {

}
