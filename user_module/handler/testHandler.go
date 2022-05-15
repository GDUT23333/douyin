package handler

import (
	"context"
	"user_module/micro_service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/14 17:52
 * @Description: TODO
 **/

type TestHandler struct{

}
func(t *TestHandler)TestLogin(ctx context.Context,request *micro_service.TestLoginRequest, response *micro_service.TestLoginResponse) error{
	response.Token = "123"
	return nil
}