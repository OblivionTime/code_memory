package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

// InitMysql 初始化MySQL
func InitMysql(dataSource string) {
	fmt.Println("init mysql")
	var err error
	DB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	//允许单表创建
	DB.SingularTable(true)
	//关闭sql语句日志
	DB.LogMode(false)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	DB.DB().SetConnMaxLifetime(time.Hour)
	//数据库迁移
	DB.AutoMigrate(&Article{}, &Category{}, &SubCategory{})
	fmt.Println("init mysql ok")
}
