package handler

import (
	"context"
	"relation_module/relation_rpc_service"
	"relation_module/service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/15 22:29
 * @Description: TODO
 **/
type RelationHandler struct{
	RelationService service.RelationService
}
//rpc relation action
func(r *RelationHandler)RelationAction(ctx context.Context, request *relation_rpc_service.DouyinRelationActionRequest, response*relation_rpc_service.DouyinRelationActionResponse) error{
	err := r.RelationService.Action(*request.UserId, *request.Token, *request.ToUserId, *request.ActionType)
	//allocate
	response.StatusCode = new(int32)
	response.StatusMsg = new(string)
	if err != nil{
		*response.StatusCode = -1
		*response.StatusMsg = err.Error()
		return nil
	}
	*response.StatusCode = 0
	*response.StatusMsg = "Success"
	return nil
}
//rpc relation follow list
func(r *RelationHandler)RelationFollowList(ctx context.Context, request *relation_rpc_service.DouyinRelationFollowListRequest, response *relation_rpc_service.DouyinRelationFollowListResponse) error{

}
//rpc relation follower list
func(r *RelationHandler)RelationFollowerList(ctx context.Context, request *relation_rpc_service.DouyinRelationFollowerListRequest,response *relation_rpc_service.DouyinRelationFollowerListResponse) error{

}