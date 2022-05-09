package dto

/**
 * @Author: Ember
 * @Date: 2022/5/9 12:00
 * @Description: TODO
 **/

type UserInfo struct{
	ID int64 `gorm:"column:id"`
	UserName string `gorm:"column:user_name"`
	UserPassWord string `gorm:"column:user_password"`
	UserNick string `gorm:"column:user_nick"`
}
func (u *UserInfo) TableName()string{
	return "user_info"
}
