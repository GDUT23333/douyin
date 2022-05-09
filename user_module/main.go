package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"user_module/controller"
	"user_module/dao"
)

func main(){
	//initial
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		dao.InitialDB()
		waitGroup.Done()
	}()
	r := gin.Default()
	go func() {
		InitRouter(r)
		waitGroup.Done()
	}()
	waitGroup.Wait()
	//start
	r.Run()
}

//初始化路由
func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", controller.ShowUserInfo)
	apiRouter.POST("/user/register/", controller.Registry)
	apiRouter.POST("/user/login/", controller.Login)
}


