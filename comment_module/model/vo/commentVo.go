package vo

/**
 * @Author: Ember
 * @Date: 2022/5/11 13:38
 * @Description: TODO
 **/

type CommentVo struct {
	Id         int64  `json:"id,omitempty"`
	User       UserVo   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}