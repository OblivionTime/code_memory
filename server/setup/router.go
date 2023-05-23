/*
 * @Author: your name
 * @Date: 2021-11-22 09:09:16
 * @LastEditTime: 2022-08-11 11:45:01
 * @LastEditors: solid
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Yun_Music_Back\setup\router.go
 */
package setup

import (
	"code_memory/controllers"
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 ,page not exists!",
	})
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.GetHeader("Origin")
		context.Header("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

//设置路由
func setupRouter(router *gin.Engine) {
	//跨域
	router.Use(Cors())
	//未找到路由
	router.NoRoute(NotFound)
	_ = mime.AddExtensionType(".js", "application/javascript")
	router.Static("/static", "./web/static")
	router.LoadHTMLFiles("./web/index.html")
	router.GET("/", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "index.html", nil)
	})
	/*
	 用户
	*/
	codeRouter := router.Group("/v1/code_memory/api")
	{
		//添加栏目
		codeRouter.POST("/add_category", controllers.CreateCategory)
		//添加文章
		codeRouter.POST("/add_article", controllers.CreateOrUpdateArticle)
		//修改文章
		codeRouter.POST("/update_article", controllers.CreateOrUpdateArticle)
		//删除文章
		codeRouter.GET("/delete_article", controllers.DeleteArticle)
		//获取文章列表
		codeRouter.GET("/list_article", controllers.GetArticleList)
		//获取所有信息
		codeRouter.GET("/info", controllers.GetInfo)
	}
}
