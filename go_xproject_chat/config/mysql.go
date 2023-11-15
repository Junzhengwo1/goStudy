package config

import (
	"fmt"
	"goStudy/go_xproject_chat/model/pojo"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MySQL = "root:12345678@tcp(127.0.0.1:3306)/go_chat?charset=utf8"
)

var (
	Db                *gorm.DB
	err               error
	datetimePrecision = 2
)

func init() {
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       MySQL,              // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{})
	fmt.Println(err, Db)
	if err != nil {
		panic(err)
	}
	if Db.Error != nil {
		panic(Db.Error)
	}

	// 执行建表语句
	err := Db.AutoMigrate(&pojo.User{})
	if err != nil {
		return
	}

	//Db.Create(&pojo.User{
	//	Username: "king",
	//	Password: "123456",
	//	Email:    "admin@admin",
	//	Phone:    "123456",
	//})

}
