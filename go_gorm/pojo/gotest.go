package pojo

type GoTest struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 必须实现 该方法

func (GoTest) TableName() string {
	return "go_test"
}
