package main

import (
	"code_memory/core"
	"code_memory/global"
	"code_memory/initialize"
)

func main() {
	//初始化配置信息
	global.CONFIG = core.Config()
	//初始化数据库
	global.DB = initialize.Gorm()
	initialize.RegisterTables() // 初始化表
	// 程序结束前关闭数据库链接
	db, _ := global.DB.DB()
	defer db.Close()
	core.RunWindowsServer()
}
