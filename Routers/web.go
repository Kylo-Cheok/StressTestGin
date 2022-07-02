package routers

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/5/14 13:14
 * @Desc:   router file
 */

import (
	setting "ginDemoProject/Pkg"
	"ginDemoProject/Routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	//路由分组
	userGroup1 := router.Group("/user")
	//userGroup1.Use(jwt.JWT())
	{
		//调用方法
		userGroup1.GET("/get_user/:username", api.GetUser)
		userGroup1.POST("/create_user", api.CreateUser)
	}

	userGroup2 := router.Group("/stress")
	//userGroup2.Use(jwt.JWT())
	{
		//调用方法
		userGroup2.POST("/start", api.Start)
	}

	return router
}
