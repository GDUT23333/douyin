package main

import (
	"github.com/gin-gonic/gin"
	"user_module/controller"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 13:48
 * @Description: TODO
 **/

//初始化路由
func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", controller.ShowUserInfo)
	apiRouter.POST("/user/register/", controller.Registry)
	apiRouter.POST("/user/login/", controller.Login)
}