package initialize

import (
	"code_memory/global"
	r "code_memory/router"
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	codeMemoryRouter := r.RouterGroupApp.System
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
	APIGroup := router.Group(global.CONFIG.System.RouterPrefix)
	{
		codeMemoryRouter.InitArticleRouter(APIGroup)
		codeMemoryRouter.InitInfoRouter(APIGroup)
		codeMemoryRouter.InitCategoryRouter(APIGroup)
	}
	return router
}
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
