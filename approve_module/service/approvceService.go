package service

import (
	"approve_module/dao"
	"approve_module/model/dto"
	"approve_module/utils"
	"errors"
	"fmt"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 10:41
 * @Description: TODO
 **/

type ApproveService interface{
	//approve action
	Action(token string,id int64,videoId int64,actionType int32) error
}
type ApproveServiceImpl struct{
	approveDao dao.ApproveDao
}
func (a *ApproveServiceImpl) Action(token string,id int64,videoId int64,actionType int32) (err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("verify token failed:",err.Error())
		return errors.New("token failed..")
	}
	//package Action
	videoAction := &dto.VideoAction{
		UserID: id,
		ActionType: actionType,
		VideoID: videoId,
	}
	//dao service
	switch actionType {
	//approve
	case 1:{
		//todo update memory
		a.approveDao.ApproveAction(videoAction)
	}
	//cancel approve
	case 2:{
		//todo update memory
		a.approveDao.CanCelApproveAction(videoAction)
	}
	}
	return
}
//single create
var(
	approveService ApproveService
	approveServiceOnce sync.Once
)
func GetApproveService() ApproveService{
	approveServiceOnce.Do(func() {
		approveService = &ApproveServiceImpl{
			approveDao: dao.GetApproveDao(),
		}
	})
	return approveService
}