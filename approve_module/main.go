package main

import (
	"approve_module/approve_rpc_service"
	"approve_module/controller"
	"approve_module/dao"
	ser "approve_module/service"
	"approve_module/handler"
	"github.com/asim/go-micro/plugins/registry/zookeeper/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
)
const(
	ServerName = "douyin.approve"
)
/**
 * @Author: Ember
 * @Date: 2022/5/11 10:07
 * @Description: TODO
 **/

func main(){
	dao.InitialDB()
	//micro registry
	newRegistry := zookeeper.NewRegistry(
		registry.Addrs("192.168.160.132:2181"))
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Address(":8083"),
		micro.Registry(newRegistry),
	)
	service.Init()
	//register user base service
	err := approve_rpc_service.RegisterApproveBaseServiceHandler(service.Server(),&handler.ApproveHandler{
		ApproveService: ser.GetApproveService(),
			})
	if err != nil{
		logger.Fatal(err)
	}
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

	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
}