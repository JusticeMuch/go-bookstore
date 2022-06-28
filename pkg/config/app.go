package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	connection_string := "justin:humannature@tcp(127.0.0.1:3306)/bookstore?charset=utf8&loc=Local&parseTime=true"
	d, err := gorm.Open("mysql", connection_string)
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
