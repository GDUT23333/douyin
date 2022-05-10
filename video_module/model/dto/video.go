package dto

import "time"

/**
 * @Author: Ember
 * @Date: 2022/5/10 10:51
 * @Description: TODO
 **/

type Video struct{
	//主键
	ID int64 `gorm:"column:id"`
	//资源URL
	PlayUrl string `gorm:"column:play_url"`
	//封面URL
	CoverUrl string `gorm:"column:cover_url"`
	//点赞数
	ApproveCount int64 `gorm:"column:approve_count"`
	//评论数
	CommentCount int64 `gorm:"column:comment_count"`
	//上传者ID
	UserId string `gorm:"column:user_id"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time"`
}
//div table name
func (v *Video) TableName() string{
	return "video"
}
