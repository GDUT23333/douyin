package vo

/**
 * @Author: Ember
 * @Date: 2022/5/11 13:39
 * @Description: TODO
 **/
type UserVo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}