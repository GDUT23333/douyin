package main

/**
 * @Author: Ember
 * @Date: 2022/5/13 22:26
 * @Description: TODO
 **/

import (
	"gateway/comment_rpc_service"
	"flag"
	"gateway/relation_rpc_service"
	"gateway/user_rpc_service"
	"github.com/golang/glog"
	"gateway/video_rpc_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"

	service "gateway/gateway_service"
)

var (
	// the user service
	userEndpoint  = flag.String("user_endpoint", "localhost:8080", "user service address")
	videoEndpoint = flag.String("video_endpoint", "localhost:8081", "video service address")
	commentEndpoint = flag.String("comment_endpoint", "localhost:8082", "comment service address")
	approveEndpoint = flag.String("approve_endpoint", "localhost:8083", "approve service address")
	relationEndpoint = flag.String("relation_endpoint", "localhost:8084", "relation service address")
	)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	//register handler
	err := service.RegisterRpcTestLoginServiceHandlerFromEndpoint(ctx,mux,*userEndpoint,opts)
	err = user_rpc_service.RegisterUserBaseServiceHandlerFromEndpoint(ctx,mux,*userEndpoint,opts)
	err = video_rpc_service.RegisterVideoBaseServiceHandlerFromEndpoint(ctx,mux,*videoEndpoint,opts)
	err = comment_rpc_service.RegisterCommentBaseServiceHandlerFromEndpoint(ctx,mux,*commentEndpoint,opts)
	err = video_rpc_service.RegisterVideoBaseServiceHandlerFromEndpoint(ctx,mux,*videoEndpoint,opts)
	err = relation_rpc_service.RegisterRelationBaseServiceHandlerFromEndpoint(ctx,mux,*relationEndpoint,opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":9090", mux)
}

func main() {
	flag.Parse()

	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}