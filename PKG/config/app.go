package config

import (
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "pylkja:asgsagjsgte32/simplerest?charset=utf8&parseTime=True&loc=Local") // username:password/tablename?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) // failed to connect to database
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
