package main

import (
	"github.com/asim/go-micro/plugins/registry/zookeeper/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"video_module/handler"
	"video_module/controller"
	"video_module/dao"
	feeService "video_module/service"
	"video_module/video_rpc_service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 22:19
 * @Description: TODO
 **/
const(
	ServerName = "douyin.video"
)

func main(){
	//initial DB
	dao.InitialDB()
	//initial server
	//patch handler
	newRegistry := zookeeper.NewRegistry(
		registry.Addrs("192.168.160.132:2181"))
	service := micro.NewService(
			micro.Name(ServerName),
			micro.Address(":8081"),
			micro.Registry(newRegistry),
		)
	//register rpc service
	//register user base service
	err := video_rpc_service.RegisterVideoBaseServiceHandler(service.Server(),&handler.VideoHandler{
		FeedService: feeService.GetFeeService(),
	})
	if err != nil{
		logger.Fatal(err)
	}
	service.Init()
	//run
	err = service.Run()
	if err != nil{
		logger.Fatal(err)
	}
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