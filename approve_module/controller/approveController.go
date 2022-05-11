package controller

import (
	"approve_module/model"
	"approve_module/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 10:46
 * @Description: TODO
 **/
var(
	approveService service.ApproveService = service.GetApproveService()
)
//approve action
func FavoriteAction(c *gin.Context) {
	id := c.Query("user_id")
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	err := approveService.Action(token, id, videoId, actionType)

	if err != nil{
		c.JSON(http.StatusOK,model.Response{
			StatusCode: -1,
			StatusMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,model.Response{
		StatusCode: 0,
		StatusMsg: "Success",
	})
	return

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {

}