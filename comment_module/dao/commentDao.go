package dao

import (
	"comment_module/model/dto"
	"comment_module/model/vo"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 12:28
 * @Description: TODO
 **/

type CommentDao interface{
	//add comment
	CreateComment(comment *dto.Comment) error
	//del comment
	DelComment(comment *dto.Comment) error
	//get comment list by video id and user id
	GetCommentList(userId int64,videoId int64) ([]*dto.Comment,error)
}

type CommentDaoImpl struct{

}
func (*CommentDaoImpl) CreateComment(comment dto.Comment)(err error){

}
func (*CommentDaoImpl) DelComment(comment dto.Comment)(err error){

}
func (*CommentDaoImpl) GetCommentList(userId int64,videoId int64) ([]*dto.Comment,error){

}
//single create
var(
	commentDao CommentDao
	commentDaoOnce sync.Once
)
func GetCommentDao() CommentDao{
	commentDaoOnce.Do(func() {
		commentDao = &CommentDaoImpl{}
	})
	return commentDao
}
