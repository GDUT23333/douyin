package dto

/**
 * @Author: Ember
 * @Date: 2022/5/11 15:26
 * @Description: TODO
 **/
type Relation struct{
	//primary key
	ID int64
	//用户ID
	UserId int64
	//关注者ID
	FollowId int64
}