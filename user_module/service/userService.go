package service

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"user_module/dao"
	"user_module/model/dto"
	"user_module/utils"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 12:14
 * @Description: TODO
 **/

type UserService interface {
	//获取用户信息
	GetUserInfo(id int64,token string) (useId int64,nickName string,err error)
	//注册用户
	RegistryUserInfo(username string,password string) (id int64,token string,err error)
	//用户登录
	Login(userName string,password string) (userId int64,token string,err error)
}

type UserServiceImpl struct{
	userDao dao.UserDao
}

func (s *UserServiceImpl) GetUserInfo(id int64,token string) (userId int64,nickName string,err error) {
	//verifyToken
	if _,_,err := utils.VerifyToken(token) ; err != nil{
		log.Fatal(err)
		return -1,"",errors.New("token not correct")
	}
	info := s.userDao.GetUserInfoById(id)
	nickName = info.UserNick
	//生成vo
	return
}

func (s *UserServiceImpl) RegistryUserInfo(username string,password string) (id int64,token string, err error){
	//verify
	 if !utils.VerifyParams(username, password){
	 	fmt.Println("username :",username,";","password:",password," not throug verify...")
	 	return -1,"",errors.New("input is not correct...")
	 }
	encryption, encryErr := utils.Md5Encryption(password)
	if encryErr != nil{
		return -1,"",encryErr
	}
	info := &dto.UserInfo{
		UserName: username,
		UserPassWord:encryption,
		UserNick: username,
	}
	//insert record
	_, insertErr := s.userDao.InsertUserInfo(info)
	if insertErr != nil{
		return -1,"", insertErr
	}
	return info.ID,utils.GenerateToken(info),nil
}

func (s *UserServiceImpl) Login(userName string,password string) (userId int64,token string,err error){
	//search info
	info := s.userDao.GetUserInfoByName(userName)
	if info == nil{
		return -1,"",errors.New("user not exist")
	}
	if !s.verifyPassword(password,info.UserPassWord){
		return -1,"",errors.New("password not correct")
	}
	//create token then return
	return info.ID,utils.GenerateToken(info),nil
}

//校验密码
func (s *UserServiceImpl) verifyPassword(password string,origin string) bool{
	encryption, err := utils.Md5Encryption(password)
	if err != nil{
		log.Fatal(err)
		return false
	}
	return encryption == origin
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