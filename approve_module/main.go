package main

import (
	"approve_module/dao"
	"github.com/gin-gonic/gin"
	"sync"
	"approve_module/controller"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 10:07
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

	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
}