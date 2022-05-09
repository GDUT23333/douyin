package controller

import (
	"github.com/gin-gonic/gin"
	 "net/http"
	"user_module/model"
	"user_module/service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 12:22
 * @Description: TODO
 **/

var (
	userService service.UserService = service.GetUserService()
)
type User struct{
	ID int64 `json:"id"`
	NickName string `json:"name"`
	FollowCount int64 `json:"follow_count"`
	FollowerCount int64 `json:"follower_count"`
	IsFollow bool `json:"is_follow"`
}
//用户信息响应体
type UserInfoResponse struct{
	model.Response
	User User
}

//注册响应体
type UserRegistryResponse struct{
	model.Response
	UserId int64 `json:"user_id,omitempty"`
	Token string `json:"token"`
}
//登录响应体
type UserLoginResponse struct{
	model.Response
	UserId int64 `json:"user_id"`
	Token string `json:"token"`
}
//登录
func Login(ctx *gin.Context){
	userName := ctx.Query("username")
	password := ctx.Query("password")
	id, token, err := userService.Login(userName, password)
	if err != nil{
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
	}else{
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
			UserId: id,
			Token: token,
		})
	}
}
//注册
func Registry(ctx *gin.Context) {
	userName := ctx.Query("username")
	password := ctx.Query("password")
	id,err := userService.RegistryUserInfo(userName, password)
	if err != nil {
		ctx.JSON(http.StatusOK, UserRegistryResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
	}else{
		ctx.JSON(http.StatusOK,UserRegistryResponse{
			Response : model.Response{
				StatusCode : 0,
				StatusMsg : "Success",
			},
			UserId: id,
			Token : "token",
		})
	}
}
//获取用户信息
func ShowUserInfo(ctx *gin.Context) {
	id := ctx.Query("id")
	token := ctx.Query("token")
	userId,nickName, err := userService.GetUserInfo(id, token)
	if err != nil{
		ctx.JSON(http.StatusOK,&UserInfoResponse{
			Response : model.Response{
				StatusCode: -1,
				StatusMsg: err.Error(),
			},
		})
	}else{
		ctx.JSON(http.StatusOK,&UserInfoResponse{
			Response:model.Response{
				StatusCode : 0,
				StatusMsg : "Success",
			},
			User : User{
				ID : userId,
				NickName: nickName,
			},
		})
	}
}