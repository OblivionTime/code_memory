package main

import (
	"code_memory/db"
	"code_memory/setup"
)

func main() {
	db.InitMysql("root:123456@tcp(127.0.0.1:3306)/article?charset=utf8mb4&parseTime=True")
	setup.InitServer("0.0.0.0:8989")
}
