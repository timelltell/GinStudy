package database

import (
	"GinStudy/log"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)
var(
	dbmysql *gorm.DB
	serverMode=-1
)
func InitMysql(){
	dsn:=fmt.Sprintf("%v:%v@tcp(%v)/%v?charet=utf8mb&parseTime=True&loc=Local",
		viper.GetString("mysql.usr"),
		viper.GetString("mysql.pword"),
		viper.GetString("mysql.address"),
		viper.GetString("mysql.databse"))
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Logger.Fatal().Msg("init mysql failed")
	}else {
		log.Logger.Info().Msg("init mysql success")
	}

	dbmysql=db
	serverMode=viper.GetInt("server.mode")
}

func CloseMysql(){

}