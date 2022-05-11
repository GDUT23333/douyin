package dao

import (
	"approve_module/model/dto"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 10:19
 * @Description: TODO
 **/

type ApproveDao interface{
	//approve video
	ApproveAction(action *dto.VideoAction)()
	//cancel approve video
	CanCelApproveAction(action *dto.VideoAction)()
	//create record
	CreateAction(action *dto.VideoAction)()
}

type ApproveDaoImpl struct{

}
//todo
func (a *ApproveDaoImpl) ApproveAction(action *dto.VideoAction){

}
func (a *ApproveDaoImpl) CanCelApproveAction(action *dto.VideoAction){

}
func (a *ApproveDaoImpl) CreateAction(action *dto.VideoAction){

}

//single create
var(
	approveDao ApproveDao
	approveDaoOnce sync.Once
)
func GetApproveDao() ApproveDao{
	approveDaoOnce.Do(func() {
		approveDao = &ApproveDaoImpl{}
	})
	return approveDao
}


