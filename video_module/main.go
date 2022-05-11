package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net/http"

	"strings"
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
	grpcServer := grpc.NewServer()
	go func() {
		InitRouter(r)
		// 判断协议是否为http/2
		// 判断是否是grpc
		r.Use(func(ctx *gin.Context) {
			// 判断协议是否为http/2
			// 判断是否是grpc
			if ctx.Request.ProtoMajor == 2 &&
				strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/grpc") {
				// 按grpc方式来请求
				grpcServer.ServeHTTP(ctx.Writer, ctx.Request)
				// 不要再往下请求了,防止继续链式调用拦截器
				ctx.Abort()
				return
			}
			// 当作普通api
			ctx.Next()
		})
		waitGroup.Done()
	}()
	waitGroup.Wait()
	// 为http/2配置参数
	h2Handle := h2c.NewHandler(r, &http2.Server{}) // 禁用TLS加密协议
	// 配置http服务
	server := &http.Server{
		Addr: ":8081",
		Handler: h2Handle,
	}
	// 启动http服务
	server.ListenAndServe()
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