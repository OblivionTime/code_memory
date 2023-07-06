package code_memory

import (
	v1 "code_memory/api/v1"

	"github.com/gin-gonic/gin"
)

type ArticleRouter struct{}

func (a *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	articleRouter := Router.Group("article")
	articleApi := v1.ApiGroupApp.ArticleApiGroup.ArticleApi
	{
		//添加文章
		articleRouter.POST("/add_article", articleApi.CreateOrUpdateArticle)
		//修改文章
		articleRouter.POST("/update_article", articleApi.CreateOrUpdateArticle)
		//删除文章
		articleRouter.GET("/delete_article", articleApi.DeleteArticle)
		//获取文章列表
		articleRouter.GET("/list_article", articleApi.GetArticleList)
	}
	return articleRouter
}
