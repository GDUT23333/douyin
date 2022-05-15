package comment_module

import (
	"comment_module/comment_rpc_service"
	"comment_module/controller"
	"comment_module/dao"
	"comment_module/handler"
	"github.com/asim/go-micro/plugins/registry/zookeeper/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	ser "comment_module/service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 12:47
 * @Description: TODO
 **/
const(
	ServerName = "douyin_comment"
)

func main() {
	dao.InitialDB()
	//micro registry
	newRegistry := zookeeper.NewRegistry(
		registry.Addrs("192.168.160.132:2181"))
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Address(":8082"),
		micro.Registry(newRegistry),
	)
	service.Init()
	//register user base service
	err := comment_rpc_service.RegisterCommentBaseServiceHandler(service.Server(),&handler.CommentHandler{
		CommentService: ser.GetCommentService(),
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

	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)
}