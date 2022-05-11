package utils

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"video_module/douyinRpc"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 20:28
 * @Description: TODO
 **/

//校验字符串参数
func VerifyParams(strs ...string) bool{
	for _,str := range strs{
		if !verify(str){
			return false
		}
	}
	return true
}
func verify(str string) bool{
	if str == ""{
		return false
	}
	return true
}


var(
	//jwt salt value
	jwtkey = []byte("douyin")
)


//verify token
func VerifyToken(token string)(*jwt.Token, *jwt.StandardClaims, error){
	Claims := &jwt.StandardClaims{}
	t, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return t, Claims, err
}

//rpc server
//GetUserInfo
func GetUserInfo(id int64) *douyinRpc.UserResponse{
	//create conn
	conn,_ := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	//create client
	client := douyinRpc.NewRpcUserServiceClient(conn)
	//dial
	response,_ := client.GetUserInfo(context.Background(), &douyinRpc.UserRequest{Id: id})
	return response
}
