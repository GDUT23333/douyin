package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"video_module/dao"
	"video_module/controller"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 22:19
 * @Description: TODO
 **/

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
	apiRouter.GET("/feed/", controller.Feed)

	apiRouter.GET("/publish/list/", controller.PublishList)
	apiRouter.POST("/publish/action/", controller.Publish)
}