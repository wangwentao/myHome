package router

import (
	"github.com/gin-gonic/gin"
	sf "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"myHome/gin/configs"
	"myHome/gin/controllers"
	_ "myHome/gin/docs"
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

	// web
	home := router.Group("/home")
	{
		home.GET("/", homeFunc)

		home.GET("/ping", pingFunc)
	}

	// swagger handler
	router.GET("/swagger/*any", gs.WrapHandler(sf.Handler))

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

// @Summary myHome project home html template page
// @Tags home
// @Success 200 string html static html template
// @Router /home [Get]
func homeFunc(c *gin.Context) {

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

// @Summary myHome project test api
// @Tags home
// @Success 200 string json pong
// @Router /home/ping [Get]
func pingFunc(c *gin.Context) {

	c.String(http.StatusOK, "pong")
}
