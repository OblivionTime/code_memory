package initialize

import (
	"code_memory/global"
	"code_memory/model/code_memory"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()

	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		code_memory.Article{},
		code_memory.Category{},
		code_memory.SubCategory{},
	)
	if err != nil {
		fmt.Println("register table failed", err)
		os.Exit(0)
	}
	fmt.Println("register table success")
}
