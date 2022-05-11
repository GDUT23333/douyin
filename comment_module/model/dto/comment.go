package dto

/**
 * @Author: Ember
 * @Date: 2022/5/11 12:27
 * @Description: TODO
 **/
type Comment struct{
	//primary key
	ID int64
	//video key
	VideoId int64
	//user key
	UserId int64
	//content
	CommentText string
	//action type 1:publish 2、del
	ActionType int
	//创建时间
	CreateTime string
}