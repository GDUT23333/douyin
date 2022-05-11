package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"relation_module/model"
	"relation_module/model/vo"
	"relation_module/service"

)

/**
 * @Author: Ember
 * @Date: 2022/5/11 15:26
 * @Description: TODO
 **/
var(
	relationService service.RelationService = service.GetRelationService()
)
type FollowListResponse struct{
	model.Response
	FollowList []vo.UserVo
}

type FollowerListResponse struct{
	model.Response
	FollowerList []vo.UserVo
}
// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	err := relationService.Action(userId, token, toUserId, actionType)
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
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	err, users := relationService.GetUserFollowList(token, userId)
	if err != nil{
		c.JSON(http.StatusOK,model.Response{
			StatusCode : -1,
			StatusMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,FollowListResponse{
		Response : model.Response{
			StatusCode : 0,
			StatusMsg : "Success",
		},
		FollowList: users,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	err, vos := relationService.GetUserFollowerList(token, userId)
	if err != nil{
		c.JSON(http.StatusOK,model.Response{
			StatusCode : -1,
			StatusMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,FollowerListResponse{
		Response : model.Response{
			StatusCode : 0,
			StatusMsg : "Success",
		},
		FollowerList: vos,
	})
}