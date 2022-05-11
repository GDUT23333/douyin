package vo

/**
 * @Author: Ember
 * @Date: 2022/5/11 19:51
 * @Description: TODO
 **/

type UserVo struct{
	//primary key
	ID int64
	//user name
	Name string
	//user follow count
	FollowCount int64
	//user follower count
	FollowerCount int64
	//user is follow owner?
	IsFollow bool
}