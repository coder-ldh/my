package initialize

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"my/global"
	"os"
)

//初始化数据库并产生数据库全局变量
func Mysql() {
	mysql := global.GVA_CONFIG.Mysql
	var url = mysql.Username + ":" + mysql.Password + "@(" + mysql.Path + ")/" + mysql.Dbname + "?" + mysql.Config
	global.GVA_LOG.Debug("MySQL连接信息:", url)
	if db, err := gorm.Open("mysql", url); err != nil {
		global.GVA_LOG.Error("MySQL启动异常:", err)
		os.Exit(0)
	} else {
		/*DB.LogMode(true)*/
		db.SingularTable(true)
		global.GVA_DB = db
		global.GVA_DB.DB().SetMaxIdleConns(mysql.MaxIdleConns)
		global.GVA_DB.DB().SetMaxOpenConns(mysql.MaxOpenConns)
		global.GVA_DB.LogMode(mysql.LogMode)
	}
}
