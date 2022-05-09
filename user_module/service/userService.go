package service

import (
	"errors"
	"strconv"
	"sync"
	"user_module/dao"
	"user_module/model/dto"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 12:14
 * @Description: TODO
 **/

type UserService interface {
	//获取用户信息
	GetUserInfo(id string,token string) (useId int64,nickName string,err error)
	//注册用户
	RegistryUserInfo(username string,password string) (id int64,err error)
	//用户登录
	Login(userName string,password string) (userId int64,token string,err error)
}

type UserServiceImpl struct{
	userDao dao.UserDao
}
//校验token
func (s *UserServiceImpl) verifyToken(token string) bool{
	return true
}

func (s *UserServiceImpl) GetUserInfo(id string,token string) (userId int64,nickName string,err error) {
	//verifyToken
	if(!s.verifyToken(token)){
		return -1,"",errors.New("token not correct")
	}
	//str change int64
	userId, err = strconv.ParseInt(id,10,64)
	info := s.userDao.GetUserInfoById(userId)
	nickName = info.UserNick
	//生成vo
	return
}

func (s *UserServiceImpl) RegistryUserInfo(username string,password string) (id int64, err error){
	info := &dto.UserInfo{
		UserName: username,
		UserPassWord:password,
	}
	//insert record
	_, insertErr := s.userDao.InsertUserInfo(info)
	if insertErr != nil{
		return -1, insertErr
	}
	return info.ID,nil
}

func (s *UserServiceImpl) Login(userName string,password string) (userId int64,token string,err error){
	if !s.verifyToken(token) {
		return -1,"",errors.New("token not correct")
	}
	//查找Info
	info := s.userDao.GetUserInfoById(userId)
	if !s.verifyPassword(password,info.UserPassWord){
		return -1,"",errors.New("password not correct")
	}
	return info.ID,s.generateToken(info),nil
}

//校验密码
func (s *UserServiceImpl) verifyPassword(password string,origin string) bool{
	return true
}
//生成Token
func (s *UserServiceImpl) generateToken(info *dto.UserInfo) string{
	return "token"
}
var(
	userService UserService
	userServiceOnce sync.Once
)
//单例创建
func GetUserService() UserService{
	userServiceOnce.Do(func() {
		userService = &UserServiceImpl{
			userDao: dao.GetUserDao(),
		}
	})
	return userService
}