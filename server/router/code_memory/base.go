package code_memory

import (
	v1 "code_memory/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (a *BaseRouter) InitInfoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.ArticleApiGroup.BaseApi
	{
		//获取文章列表
		baseRouter.GET("/info", baseApi.GetInfo)
	}
	return baseRouter
}
