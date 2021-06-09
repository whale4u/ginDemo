package common

import (
	"fmt"
	"ginDemo/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	fmt.Println("数据库连接")
	InitDB()
}

func InitDB() *gorm.DB {
	host := config.MYSQL_HOST
	port := config.MYSQL_PORT
	database := config.MYSQL_DB
	username := config.MYSQL_USERNAME
	password := config.MYSQL_PASSWORD
	chaset := config.MYSQL_CHARSET

	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		username,
		password,
		host,
		port,
		database,
		chaset)

	fmt.Println("数据库链接：", sqlStr)

	db, err := gorm.Open(mysql.Open(sqlStr))
	if err != nil {
		fmt.Println("打开数据库失败", err)
		panic("打开数据库失败" + err.Error())
	}
	DB = db
	return DB
}
