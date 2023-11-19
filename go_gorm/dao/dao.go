package dao

import (
	"fmt"
	"goStudy/go_gorm/config"
	"goStudy/go_gorm/pojo"
)

//nil 表示空值，用于表示指针、通道、映射、切片和接口类型的空值

func Query(id int) *pojo.GoTest {
	// 查询单条数据
	var test *pojo.GoTest
	//err := conf.Db.Where("id = ?", id).First(&test).Error  等价
	err := config.Db.First(&test, id).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
	if test.Id == 0 {
		return nil
	}
	return test
}

func QueryByIds(ids []int) []pojo.GoTest {
	// 	切片
	var testList []pojo.GoTest
	// 这里没查 id 而ID在结构体是int 会被赋默认值0 注意点
	config.Db.Select("name,age").Find(&testList, ids)
	if len(testList) == 0 {
		fmt.Println("查询结果为空")
	}
	for _, test := range testList {
		fmt.Println(test)
	}
	return nil
}
