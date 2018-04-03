package models

import (
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type RTModel struct {
	gorm.Model
}
var RTDB gorm.DB
func init() {
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "rt_" + defaultTableName[0:len(defaultTableName) - 1];
	}
	RTDB, err := gorm.Open("mysql", "root:root@/rest-trains?charset=utf8")
	if err != nil {
		log.Fatal("创建数据库连接失败", err)
		return
	}
	RTDB.SetLogger(log.New(os.Stdout, "\r\n", 0))
	RTDB.LogMode(true)
	RTDB.Set("gorm:table_options DEFAULT CHARSET=utf8", "ENGINE=InnoDB")

	RTDB.AutoMigrate(&User{})
	RTDB.AutoMigrate(&Document{})
	RTDB.AutoMigrate(&Project{})
}
