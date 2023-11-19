package config

// 系统出初始化

import (
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"log"
)

type Conf struct {
	Port         string `mapstructure:"port"`
	Host         string `mapstructure:"host"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

var (
	Config Conf
)

// mapstructure

func initConfig() {
	// 读yml
	viper.SetConfigName("app")
	viper.AddConfigPath("./go_x_project_online_test/conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read conf is fail", err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalln("unmarshal conf is fail", err)
	}

}

func init() {
	initConfig() // yml 配置文件
	initMysql()

}
