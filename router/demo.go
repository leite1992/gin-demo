package router

import (
	"../apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() (engine *gin.Engine) {
	/** 1.define gin obj*/
	ginObj := gin.Default()

	/** 2.init the group router*/
	demoGroup := ginObj.Group("/demo")
	demoGroup.GET("/test", (&apis.Demo{}).GetDemo)
	demoGroup.POST("/test", (&apis.Demo{}).PostDemo)

	return ginObj
}
