package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net/http"
	"strings"
	"sync"
	"user_module/controller"
	"user_module/dao"
	"user_module/douyinRpc"
)

func main(){
	//initial
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)
	go func() {
		dao.InitialDB()
		waitGroup.Done()
	}()
	r := gin.Default()
	grpcServer := grpc.NewServer()
	go func() {
		InitRouter(r)
		// global interceptor
		r.Use(func(ctx *gin.Context) {
			// is http/2
			// is grpc
			if ctx.Request.ProtoMajor == 2 &&
				strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/grpc") {
				// 按grpc方式来请求
				grpcServer.ServeHTTP(ctx.Writer, ctx.Request)
				// 不要再往下请求了,防止继续链式调用拦截器
				ctx.Abort()
				return
			}
			//comon api
			ctx.Next()
		})
		waitGroup.Done()
	}()
	go func() {
		//initial rpc server
		douyinRpc.RegisterRpcUserServiceServer(grpcServer,&douyinRpc.UserRpcServiceImpl{})
		waitGroup.Done()
	}()
	waitGroup.Wait()

	//initial http/2
	h2Handle := h2c.NewHandler(r, &http2.Server{}) // 禁用TLS加密协议
	//set up http service
	server := &http.Server{
		Addr: ":8080",
		Handler: h2Handle,
	}
	//start
	server.ListenAndServe()

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


