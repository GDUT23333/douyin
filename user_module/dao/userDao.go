package dao

import (
	"sync"
	"user_module/model/dto"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 12:03
 * @Description: TODO
 **/

type UserDao interface{
	//根据Id获取用户信息
	GetUserInfoById(id int64) *dto.UserInfo
	//根据账号获取用户信息
	GetUserInfoByName(userName string) *dto.UserInfo
	//添加新的用户信息
	InsertUserInfo(userInfo *dto.UserInfo) (count int64,err error)
}

type UserDaoImpl struct{

}
func (u *UserDaoImpl) GetUserInfoById(id int64) *dto.UserInfo {
	var userInfo dto.UserInfo
	UserDB.First(&userInfo,id)
	return &userInfo
}
func (u *UserDaoImpl) GetUserInfoByName(userName string) *dto.UserInfo {
	var userInfo dto.UserInfo
	UserDB.Where("user_name = ?",userName).First(&userInfo)
	return &userInfo
}
func (u *UserDaoImpl) InsertUserInfo(userInfo *dto.UserInfo) (count int64 ,err error){
	result := UserDB.Create(&userInfo)
	count = result.RowsAffected
	err = result.Error
	return
}

var(
	userDao UserDao
	userDaoOnce sync.Once
)
//单例创建
func GetUserDao() UserDao{
	userDaoOnce.Do(func() {
		userDao = &UserDaoImpl{}
	})
	return userDao
}


