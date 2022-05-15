package service

import (
	"comment_module/comment_rpc_service"
	"comment_module/dao"
	"comment_module/model/dto"
	"comment_module/utils"
	"errors"
	"fmt"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 12:34
 * @Description: TODO
 **/
type CommentService interface{
	//public action
	PublicAction(userId int64,token string,videoId int64,commentText string)error
	//del action
	DelAction(userId int64,token string,videoId int64,commentId int64) error
	//comment list
	CommentList(userId int64,token string,videoId int64)(list []*comment_rpc_service.Comment,err error)
}
type CommentServiceImpl struct{
	commentDao dao.CommentDao
}
func (c *CommentServiceImpl) PublicAction(userId int64,token string,videoId int64,commentText string) (err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("verify token failed..")
		return
	}
	//check params
	flag := utils.VerifyParams(commentText)
	if !flag{
		fmt.Println("params is empty...")
		err = errors.New("params is empty..")
		return
	}
	//package
	comment := &dto.Comment{
		UserId: userId,
		VideoId: videoId,
		CommentText: commentText,
		ActionType: 1,
	}
	err = c.commentDao.CreateComment(comment)
	return
}
func (c *CommentServiceImpl) DelAction(userId int64,token string,videoId int64,commentId int64) (err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("verify token failed..")
		return
	}
	//package data
	comment := &dto.Comment{
		UserId: userId,
		VideoId: videoId,
		ActionType: 1,
	}
	err = c.commentDao.DelComment(comment)
	return
}
func (c *CommentServiceImpl) CommentList(userId int64,token string,videoId int64)(list []*comment_rpc_service.Comment,err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("verify token failed..")
		return
	}

	commentList, err := c.commentDao.GetCommentList(userId, videoId)
	//dto change vo
	commentVoList := make([]*comment_rpc_service.Comment,len(commentList))
	for _,comment := range(commentList){
		commentVoList = append(commentVoList,&comment_rpc_service.Comment{
			Id:&comment.ID,
			Content: &comment.CommentText,
			CreateDate: &comment.CreateTime,
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