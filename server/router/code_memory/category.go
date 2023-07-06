package code_memory

import (
	v1 "code_memory/api/v1"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

func (a *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	categoryRouter := Router.Group("category")
	categoryApi := v1.ApiGroupApp.ArticleApiGroup.CategoryApi
	{
		//添加栏目
		categoryRouter.POST("/add_category", categoryApi.CreateCategory)
	}
	return categoryRouter
}
