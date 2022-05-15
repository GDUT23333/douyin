package handler

import (
	"context"
	"user_module/service"
	"user_module/user_rpc_service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/12 22:21
 * @Description: TODO
 **/
type UserHandler struct{
	UserService service.UserService
}

//login service
func(u *UserHandler)Login(ctx context.Context,request *user_rpc_service.DouyinUserLoginRequest, response *user_rpc_service.DouyinUserLoginResponse)  error{
	id, token, err := u.UserService.Login(request.Username, request.Password)
	if err != nil{
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		response.Token = token
		response.UserId = id
		return nil
	}
	response.StatusCode = 0
	response.StatusMsg = "Success"
	response.Token = token
	response.UserId = id
	return nil
}
//register service
func(u *UserHandler)Register(ctx context.Context, request *user_rpc_service.DouyinUserRegisterRequest,response *user_rpc_service.DouyinUserLoginResponse)  error{
	id, token, err := u.UserService.RegistryUserInfo(request.Username, request.Password)
	if err != nil{
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return nil
	}
	response.StatusCode = 0
	response.StatusMsg = "Success"
	response.Token = token
	response.UserId = id
	return nil
}
//show user info service
func(u *UserHandler)ShowUserInfo(ctx context.Context,request *user_rpc_service.DouyinUserRequest,response *user_rpc_service.DouyinUserResponse)  error{
	id, name, err := u.UserService.GetUserInfo(request.UserId, request.Token)
	if err != nil{
		response.StatusCode = -1;
		response.StatusMsg = err.Error()
		return nil
	}
	response.StatusCode = 0;
	response.StatusMsg = "Success"
	response.User = &user_rpc_service.User{
		Id: id,
		Name: name,
		//TODO follow waitting...
	}
	return nil

}
