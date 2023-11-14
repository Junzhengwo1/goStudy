package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

const (
	Type  = "mysql"
	MySQL = "root:12345678@tcp(127.0.0.1:3306)/for_go?charset=utf8"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(Type, MySQL)
	fmt.Println(err, Db)
	if err != nil {
		panic(err)
	}
	if Db.Error != nil {
		panic(Db.Error)
	}
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
}
