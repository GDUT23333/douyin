package relation_module

import (
	"github.com/gin-gonic/gin"
	"relation_module/controller"
	"relation_module/dao"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 15:24
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

	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}