package main

import (
	"github.com/asim/go-micro/plugins/registry/zookeeper/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"user_module/dao"
	hler "user_module/handler"
	"user_module/micro_service"
	userService "user_module/service"
	"user_module/user_rpc_service"
)
const(
	ServerName = "douyin.user"
)
func main(){
	//initial DB
	dao.InitialDB()
	//micro registry
	newRegistry := zookeeper.NewRegistry(
		registry.Addrs("192.168.160.132:2181"))
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Address(":8080"),
		micro.Registry(newRegistry),
	)
	service.Init()
	//register user base service
	err := user_rpc_service.RegisterUserBaseServiceHandler(service.Server(),&hler.UserHandler{
		UserService: userService.GetUserService(),
	})
	micro_service.RegisterRpcTestLoginServiceHandler(service.Server(),&hler.TestHandler{})
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
//func InitRouter(r *gin.Engine) {
//	// public directory is used to serve static resources
//	r.Static("/static", "./public")
//
//	apiRouter := r.Group("/douyin")
//
//	apiRouter.GET("/user/", controller.ShowUserInfo)
//	apiRouter.POST("/user/register/", controller.Registry)
//	apiRouter.POST("/user/login/", controller.Login)
//}


