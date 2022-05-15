package handler

import (
	"approve_module/approve_rpc_service"
	"approve_module/service"
	"context"
)

/**
 * @Author: Ember
 * @Date: 2022/5/15 17:58
 * @Description: TODO
 **/
type ApproveHandler struct{
	ApproveService service.ApproveService
}
func(a *ApproveHandler)FavoriteAction(ctx context.Context, request *approve_rpc_service.DouyinFavoriteActionRequest, response*approve_rpc_service.DouyinFavoriteActionResponse) error{
	err := a.ApproveService.Action(*request.Token, *request.UserId, *request.VideoId, *request.ActionType)
	//allocate
	response.StatusCode = new(int32)
	response.StatusMsg = new(string)
	if err != nil{
		*response.StatusCode = -1
		*response.StatusMsg = err.Error()
		return nil
	}
	*response.StatusMsg = "Success"
	*response.StatusCode = 0
	return nil
}

func(a *ApproveHandler)FavoriteList(ctx context.Context, request *approve_rpc_service.DouyinFavoriteListRequest, response *approve_rpc_service.DouyinFavoriteListResponse) error{
	//todo
	return nil
}