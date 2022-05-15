package controller

import (
	"github.com/gin-gonic/gin"
	"video_module/model"
	"video_module/model/vo"
	"video_module/service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/10 10:49
 * @Description: TODO
 **/
var (
	feeService service.FeeService = service.GetFeeService()
)
//public list response
type PublishListResponse struct{
	model.Response
	VideoList []vo.VideoVo `json:"video_list"`
}


//introduce feed
func Feed(c *gin.Context){

}

//publish feed
func Publish(c * gin.Context){

}

//publish user feed list
func PublishList(c *gin.Context){

}
