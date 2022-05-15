package handler

import (
	"comment_module/comment_rpc_service"
	"comment_module/service"
	"context"
)

/**
 * @Author: Ember
 * @Date: 2022/5/15 16:49
 * @Description: TODO
 **/

type CommentHandler struct{
	CommentService service.CommentService
}
func (c *CommentHandler)CommentAction(ctx context.Context, request *comment_rpc_service.DouyinCommentActionRequest, response*comment_rpc_service.DouyinCommentActionResponse) error{
	actionType := *request.ActionType
	var err error
	switch actionType {
	case 1:{
		err = c.CommentService.PublicAction(*request.UserId, *request.Token, *request.VideoId, *request.CommentText)
	}
	case 2:{
		err = c.CommentService.DelAction(*request.UserId, *request.Token, *request.VideoId, *request.CommentId)
	}
	}
	//allocate
	response.StatusMsg = new(string)
	response.StatusCode = new(int32)
	if err != nil{
		*response.StatusMsg = err.Error()
		*response.StatusCode = -1
		return nil
	}
	*response.StatusMsg = "Success"
	*response.StatusCode = 0
	return nil
}
func (c *CommentHandler)CommentList(ctx context.Context,request *comment_rpc_service.DouyinCommentListRequest, response *comment_rpc_service.DouyinCommentListResponse) error{
	list, err := c.CommentService.CommentList(*request.UserId, *request.Token, *request.VideoId)
	//allocate
	response.StatusMsg = new(string)
	response.StatusCode = new(int32)
	if err != nil{
		*response.StatusMsg = err.Error()
		*response.StatusCode = -1
		return nil
	}
	*response.StatusMsg = "Success"
	*response.StatusCode = 0
	response.CommentList = list
	return nil
}