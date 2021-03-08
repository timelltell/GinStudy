package database

import (
	"GinStudy/log"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var(
	dbmysql *gorm.DB
	serverMode=-1
)
func InitMysql(){
	dsn:=fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
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

func DB(tableName string) *gorm.DB{

	return dbmysql.Table(tableName)
}
type PlanVersion struct {
	Id int `json:"-" gorm:"column:id;primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}
func Createtable(){
	dbmysql.AutoMigrate(&PlanVersion{})

	return
}