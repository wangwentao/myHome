package router

import (
	"github.com/gin-gonic/gin"
	"myHome/gin/configs"
	"myHome/gin/controllers"
	"net/http"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(configs.Server.RunMode)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// miniprogram
	mini := router.Group("/mini")
	{
		mini.GET("/login", controllers.MiniLogin)
		mini.POST("/profile", controllers.UserProfile)
	}

	//router.GET("/mini/login", controllers.MiniLogin)
	/*router.GET("/mini/login", controllers.MiniLogin)
	router.POST("/mini/profile", controllers.UserProfile)*/

	// web
	home := router.Group("/home")
	{
		home.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Main website",
			})
		})

		home.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	/*// html
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})*/

	// wechat
	/*router.GET("/wechat", controllers.ServeWeChat)
	router.POST("/wechat", controllers.ServeWeChat)*/

	return router
}
