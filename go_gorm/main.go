package main

import (
	"fmt"
	"goStudy/go_gorm/dao"
)

func main() {

	query := dao.Query(1)
	fmt.Println(*query)
	dao.QueryByIds([]int{1, 2})

}
