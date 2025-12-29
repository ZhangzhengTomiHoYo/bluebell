package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录用户，判断请求头中是否有 有效的JWT
		c.String(http.StatusOK, "你是登录用户，你进来了")
	})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "no Route!!! 404",
		})
	})
	return r
}
