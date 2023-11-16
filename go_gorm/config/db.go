package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

const (
	MySQL  = "root:12345678@tcp(127.0.0.1:3306)/go_for?charset=utf8"
	MySQL2 = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
)

var (
	Db                *gorm.DB
	err               error
	datetimePrecision = 2
)

func init() {

	root := fmt.Sprintf(MySQL2, "root", "12345678", "127.0.0.1", 3306, "go_for")

	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       root,               // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	fmt.Println(err, Db)
	if err != nil {
		log.Fatalln("DB connect fail", err)
	}
	if Db.Error != nil {
		log.Fatalln("DB connect fail", err)
	}
}
