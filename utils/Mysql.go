package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitMysql() *gorm.DB {
	var err error

	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gocron?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "go_" + defaultTableName
	}

	DB.DB().SetMaxOpenConns(10)
	DB.DB().SetMaxIdleConns(10)
	DB.SingularTable(true)
	return DB
}
