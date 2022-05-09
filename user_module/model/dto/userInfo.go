package dto

/**
 * @Author: Ember
 * @Date: 2022/5/9 12:00
 * @Description: TODO
 **/

type UserInfo struct{
	ID int64 `grom:"column:id"`
	UserName string `grom:"column:user_name"`
	UserPassWord string `grom:"column:user_password"`
	UserNick string `grom:"column:user_nick"`
}
func (u *UserInfo) TableName()string{
	return "user_info"
}
