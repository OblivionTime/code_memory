package core

import (
	"code_memory/global"
	"code_memory/initialize"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func RunWindowsServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	fmt.Println(s.ListenAndServe().Error())
}
