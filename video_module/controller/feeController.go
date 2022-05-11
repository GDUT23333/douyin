package controller

import (
	"github.com/gin-gonic/gin"
	"video_module/model/vo"
	"video_module/service"
	"video_module/model"
	"net/http"
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
	token := c.Query("token")
	id := c.Query("user_id")
	data, err := c.FormFile("data")
	if err != nil{
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	_, err = feeService.PublishFee(token, id, data, c)
	if err != nil{
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg: "uploaded successfully",
	})
}

//publish user feed list
func PublishList(c *gin.Context){
	token := c.Query("token")
	id := c.Query("user_id")
	publishFees, err := feeService.GetPublishFees(id, token)
	if err != nil{
		c.JSON(http.StatusOK,PublishListResponse{
			Response : model.Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK,PublishListResponse{
		Response : model.Response{
			StatusCode: 0,
			StatusMsg: "Success",
		},
		VideoList: publishFees,
	})
	return
}
