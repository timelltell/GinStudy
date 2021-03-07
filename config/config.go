package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig(config_file string) {
	fmt.Printf("config file is %s\n", config_file)
	viper.SetConfigFile(config_file)
	viper.SetConfigType("yaml")
	err:=viper.ReadInConfig()
	if err != nil{
		fmt.Println("init config error")
	}

}