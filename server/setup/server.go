package setup

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitServer(host string) {
	router := gin.New()
	// router := gin.Default()
	//设置路由
	setupRouter(router)
	err := router.Run(host)
	if err != nil {
		fmt.Println("Init http server. Error :", err)
	}
}
