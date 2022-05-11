package controller

import (
	"comment_module/model"
	"comment_module/model/vo"
	"comment_module/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 12:39
 * @Description: TODO
 **/

var(
	commentService service.CommentService = service.GetCommentService()
)
//comment list response
type CommentListResponse struct{
	model.Response
	CommentList []vo.CommentVo
}

//comment action
func CommentAction(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	//choose service by actionType
	var err error
	switch actionType{
	case "1":{
		commentText := c.Query("comment_text")
		err = commentService.PublicAction(userId, token, videoId, commentText)
	}
	case "2":{
		commentId := c.Query("comment_id")
		err = commentService.DelAction(userId, token, videoId, commentId)
	}
	}
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

//comment list
func CommentList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	videoId := c.Query("video_id")
	commentList, err := commentService.CommentList(userId, token, videoId)
	if err != nil{
		c.JSON(http.StatusOK,model.Response{
			StatusCode: -1,
			StatusMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,CommentListResponse{
		Response : model.Response{
			StatusCode: 0,
			StatusMsg: "Success",
		},
		CommentList: commentList,
	})
}