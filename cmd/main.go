package main

import (
	"GinStudy/config"
	"GinStudy/database"
	"GinStudy/handler"
	"GinStudy/log"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"

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
	handler.Ginserver=gin.New()
	handler.InitService()

	srv:=&http.Server{
		Addr: viper.GetString("server.host")+":"+viper.GetString("server.port"),
		Handler: handler.Ginserver,
	}
	//log.Logger.Info().Msgf("address  %v",srv.Addr)
	//log.Logger.Info().Msg(viper.GetString("server.host"))
	//log.Logger.Info().Msg(viper.GetString("server.port"))
	//log.Logger.Info().Msg(viper.GetString("server.port"))
	//log.Logger.Info().Msg(srv.Addr)
	//log.Logger.Info().Msgf("address  %v",srv.Addr)

	func(){
		if err:=srv.ListenAndServe();err!=nil && err!=http.ErrServerClosed{
			log.Logger.Error().Msg("server closed with error")
		}
	}()

}
