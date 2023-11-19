package config

// 系统出初始化

import (
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
)

type Conf struct {
	Port         string `mapstructure:"port"`
	Host         string `mapstructure:"host"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

var (
	Config    Conf
	configErr error
)

// mapstructure

func initConfig() {
	// 读yml
	viper.SetConfigName("app")
	viper.AddConfigPath("./go_x_project_online_test/conf")
	configErr = viper.ReadInConfig()
	if configErr != nil {
		panic(configErr)
	}
	configErr = viper.Unmarshal(&Config)
	if configErr != nil {
		panic(configErr)
	}
}

func Init() {
	initConfig() // yml 配置文件
	initMysql()

}
