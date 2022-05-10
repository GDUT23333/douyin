package vo

/**
 * @Author: Ember
 * @Date: 2022/5/10 15:02
 * @Description: TODO
 **/

type UserVo struct{
	//主键
	ID int64
	//名字
	Name string
	//关注数
	FollowCount int64
	//粉丝数
	FollowerCount int64
	//是否关注
	IsFollow bool
}