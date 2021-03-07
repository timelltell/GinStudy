package main

import (
	"fmt"
	"GinStudy/config"
	"GinStudy/log"
	"GinStudy/database"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	//"github.com/gin-gonic/gin"
	//"html/template"
	//"log"
	//"net/http"
)

var(
	config_file=flag.String("config","./config/config.yaml","input path of config file")
)
func main() {
	fmt.Println("----------------- main start -----------------")
	//初始化入参
	flag.Parse()
	//初始化配置
	config.InitConfig(*config_file)
	//初始化日志
	log.InitLog()
	//初始化数据库
	database.InitMysql()
	defer database.CloseMysql()
	//初始化server

}
