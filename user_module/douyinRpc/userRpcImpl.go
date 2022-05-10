package douyinRpc

import (
	"context"
	"user_module/dao"
)

/**
 * @Author: Ember
 * @Date: 2022/5/10 22:08
 * @Description: TODO
 **/

type UserRpcServiceImpl struct{
	userDao dao.UserDao
}
func (u *UserRpcServiceImpl) GetUserInfo(ctx context.Context,request *UserRequest) (*UserResponse, error){
	id := request.Id
	info := u.userDao.GetUserInfoById(id)
	return &UserResponse{
		Id: info.ID,
		Name: info.UserNick,
		//todo followCount and followerCount ?
	},nil
}