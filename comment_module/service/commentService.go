package service

import (
	"comment_module/dao"
	"comment_module/model/dto"
	"comment_module/model/vo"
	"comment_module/utils"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 12:34
 * @Description: TODO
 **/
type CommentService interface{
	//public action
	PublicAction(userId string,token string,videoId string,commentText string)error
	//del action
	DelAction(userId string,token string,videoId string,commentId string) error
	//comment list
	CommentList(userId string,token string,videoId string)(list []vo.CommentVo,err error)
}
type CommentServiceImpl struct{
	commentDao dao.CommentDao
}
func (c *CommentServiceImpl) PublicAction(userId string,token string,videoId string,commentText string) (err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("verify token failed..")
		return
	}
	//check params
	flag := utils.VerifyParams(userId, videoId, commentText)
	if !flag{
		fmt.Println("params is empty...")
		err = errors.New("params is empty..")
		return
	}
	//change params
	uId,err := strconv.ParseInt(userId,10,64)
	if err != nil{
		fmt.Println("user id change int64 failed:",err.Error())
		return
	}
	vId ,err := strconv.ParseInt(videoId,10,64)
	if err != nil{
		fmt.Println("video id change int64 failed:",err.Error())
		return
	}
	//package
	comment := &dto.Comment{
		UserId: uId,
		VideoId: vId,
		CommentText: commentText,
		ActionType: 1,
	}
	err = c.commentDao.CreateComment(comment)
	return
}
func (c *CommentServiceImpl) DelAction(userId string,token string,videoId string,commentId string) (err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("verify token failed..")
		return
	}
	//check params
	flag := utils.VerifyParams(userId, videoId, commentId)
	if !flag{
		fmt.Println("params is empty...")
		err = errors.New("params is empty..")
		return
	}
	//change params
	uId,err := strconv.ParseInt(userId,10,64)
	if err != nil{
		fmt.Println("user id change int64 failed:",err.Error())
		return
	}
	vId ,err := strconv.ParseInt(videoId,10,64)
	if err != nil{
		fmt.Println("video id change int64 failed:",err.Error())
		return
	}
	//package data
	comment := &dto.Comment{
		UserId: uId,
		VideoId: vId,
		ActionType: 1,
	}
	err = c.commentDao.DelComment(comment)
	return
}
func (c *CommentServiceImpl) CommentList(userId string,token string,videoId string)(list []vo.CommentVo,err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("verify token failed..")
		return
	}
	//check params
	flag := utils.VerifyParams(userId, videoId)
	if !flag{
		fmt.Println("params is empty...")
		err = errors.New("params is empty..")
		return
	}
	//change params
	uId,err := strconv.ParseInt(userId,10,64)
	if err != nil{
		fmt.Println("user id change int64 failed:",err.Error())
		return
	}
	vId ,err := strconv.ParseInt(videoId,10,64)
	if err != nil{
		fmt.Println("video id change int64 failed:",err.Error())
		return
	}
	commentList, err := c.commentDao.GetCommentList(uId, vId)
	//dto change vo
	commentVoList := make([]vo.CommentVo,len(commentList))
	for _,comment := range(commentList){
		commentVoList = append(commentVoList,vo.CommentVo{
			Id:comment.ID,
			Content: comment.CommentText,
			CreateDate: comment.CreateTime,
		})
	}
	//todo rpc patch get userinfo
	return commentVoList,err
}
//single create
var(
	commentService CommentService
	commentServiceOnce sync.Once
)
func GetCommentService() CommentService{
	commentServiceOnce.Do(func() {
		commentService = &CommentServiceImpl{
			commentDao: dao.GetCommentDao(),
		}
	})
	return commentService
}