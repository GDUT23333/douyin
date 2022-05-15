package relation_module

import (
	"github.com/asim/go-micro/plugins/registry/zookeeper/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"relation_module/controller"
	"relation_module/dao"
	"relation_module/relation_rpc_service"
	"sync"
	"relation_module/handler"
	ser "relation_module/service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 15:24
 * @Description: TODO
 **/
const(
	ServerName = "douyin.relation"
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
		micro.Address(":8084"),
		micro.Registry(newRegistry),
	)
	//register rpc service
	//register base service
	err := relation_rpc_service.RegisterRelationBaseServiceHandler(service.Server(),&handler.RelationHandler{
		RelationService: ser.GetRelationService(),
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

	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}