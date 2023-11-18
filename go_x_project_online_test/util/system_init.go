package util

// 系统出初始化

import (
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"log"
)

var (
	mysqlConf  map[string]interface{}
	ServerConf map[string]interface{} // yml 配置信息
)

func initConfig() {
	// 读yml
	viper.SetConfigName("app")
	viper.AddConfigPath("./go_x_project_online_test/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read config is fail", err)
	}
	mysqlConf = viper.Get("mysql").(map[string]interface{})   // 类型断言
	ServerConf = viper.Get("server").(map[string]interface{}) // 类型断言
	log.Printf("Using mysql file type is:%T \nvalue is:%v", mysqlConf, mysqlConf)
	log.Printf("Using server file type is:%T \nvalue is:%v", ServerConf, ServerConf)
}

func init() {
	initConfig() // yml 配置文件
	initMysql()

}
