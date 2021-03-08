package model

import (
	"GinStudy/database"
	//"context"
)

type PlanVersion struct {
	Id int `json:"-" gorm:"column:id;primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}
func GetName() (string,error){
	p:=PlanVersion{}

	database.Createtable()

	tx:=database.DB("test").
		Where("id=?",1).
		Find(&p)
	return p.Name, tx.Error
}
