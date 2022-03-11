package initRouter

import (
	"diy_ginHello/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/statics", "./statics")

	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/register",handler.UserRegister)
		userGroup.POST("/login",handler.UserLogin)
		userGroup.GET("/profile/",handler.UserProfile)
		userGroup.POST("/update",handler.UpdateUserProfile)
	}
	articleGroup := router.Group("")
	{
		articleGroup.GET("/article/:id",handler.GetOne)
		articleGroup.GET("/articles",handler.GetAll)
		articleGroup.POST("/article",handler.Insert)
		articleGroup.DELETE("/article/:id",handler.DeleteOne)
	}

	return router
}