package handler

import (
	"context"
	"video_module/service"
	"video_module/video_rpc_service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/14 21:07
 * @Description: TODO
 **/

type VideoHandler struct{
	FeedService service.FeeService
}
//push feed
func (v *VideoHandler)PushFeed(ctx context.Context, request *video_rpc_service.DouyinFeedRequest, response *video_rpc_service.DouyinFeedResponse) error{
	feed, nextTime, err := v.FeedService.PushFeed(*request.LatestTime)
	if err != nil{
		*response.StatusCode = -1
		*response.StatusMsg = err.Error()
		return nil
	}
	*response.StatusCode = 0
	*response.StatusMsg = "Success"
	response.VideoList = feed
	*response.NextTime = nextTime
	return nil
}
//publish video
func (v *VideoHandler)PublishFee(ctx context.Context,request *video_rpc_service.DouyinPublishActionRequest, response *video_rpc_service.DouyinPublishActionResponse) error{
	//TODO bytes -> FileHeader
	_, err := v.FeedService.PublishFee(*request.Token, *request.UserId, request.Data)
	if err != nil{
		*response.StatusCode = -1
		*response.StatusMsg = err.Error()
		return nil
	}
	*response.StatusCode = 0
	*response.StatusMsg = "Success"
	return nil
}
//get fees
func (v *VideoHandler)GetPublishFees(ctx context.Context, request *video_rpc_service.DouyinPublishListRequest,response  *video_rpc_service.DouyinPublishListResponse) error{
	fees, err := v.FeedService.GetPublishFees(*request.UserId,*request.Token)
	response.StatusCode = new(int32)
	response.StatusMsg = new(string)
	if err != nil{
		*response.StatusCode = -1
		*response.StatusMsg = err.Error()
		return nil
	}
	*response.StatusCode = 0
	*response.StatusMsg = "Success"
	response.VideoList = fees
	return nil
}