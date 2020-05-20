package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

/*初始化时运行的方法*/
func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/fantasy?charset=utf8&parseTime=True&readTimeout=500ms")
	if err != nil {
		log.Printf("mysql connect error %v\n", err)
	}
	if DB.Error != nil {
		log.Printf("database error %v\n", DB.Error)
	}
	DB.LogMode(true)
	DB.SingularTable(true)
}
