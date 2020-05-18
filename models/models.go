package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

var name *string

func init(){
	var err error
	println(DB)
	println(name)

	println(name)
	DB, err := gorm.Open("mysql", "root:12345678@tcp(localhost:3306)/book?charset=utf8&parseTime=True&readTimeout=500ms")
	if err != nil {
		log.Printf("mysql connect error %v\n", err)
	}
	if DB.Error != nil {
		log.Printf("database error %v\n", DB.Error)
	}
	DB.SingularTable(true)
	println(DB)
}
