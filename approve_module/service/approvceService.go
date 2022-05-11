package service

import (
	"approve_module/dao"
	"approve_module/model/dto"
	"approve_module/utils"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 10:41
 * @Description: TODO
 **/

type ApproveService interface{
	//approve action
	Action(token string,id string,videoId string,actionType string) error
}
type ApproveServiceImpl struct{
	approveDao dao.ApproveDao
}
func (a *ApproveServiceImpl) Action(token string,id string,videoId string,actionType string) (err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("verify token failed:",err.Error())
		return errors.New("token failed..")
	}
	//check params
	verifuResult := utils.VerifyParams(id, videoId, actionType)
	if !verifuResult{
		return errors.New("params is not correct,may be empty")
	}
	//change params
	userId,err := strconv.ParseInt(id,10,64)
	if err != nil{
		fmt.Println("user id changes failed:",err.Error())
		return
	}
	vId,err := strconv.ParseInt(videoId,10,64)
	if err != nil{
		fmt.Println("video id changes failed:",err.Error())
		return
	}
	atype,err := strconv.Atoi(actionType)
	if err != nil{
		fmt.Println("action type changes failed:",err.Error())
		return
	}
	//package Action
	videoAction := &dto.VideoAction{
		UserID: userId,
		ActionType: atype,
		VideoID: vId,
	}
	//dao service
	switch atype {
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