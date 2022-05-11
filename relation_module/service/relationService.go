package service

import (
	"errors"
	"fmt"
	"relation_module/dao"
	"relation_module/model/dto"
	"relation_module/model/vo"
	"relation_module/utils"
	"strconv"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 15:26
 * @Description: TODO
 **/

type RelationService interface{
	// relation action
	// actionType = 1 => action
	// actionType = 2 => cancel action
	Action(userId string,token string,toUserId string,actionType string) error
	//get user follow list
	GetUserFollowList(token string,userId string)(error,[]vo.UserVo)
	//get user follower list
	GetUserFollowerList(token string,userId string)(error,[]vo.UserVo)
}
type RelationServiceImpl struct{
	relationDao dao.RelationDao
}
func (r *RelationServiceImpl)Action(userId string,token string,toUserId string,actionType string) (err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("token verify failed:",err.Error())
		return
	}
	//check params
	if !utils.VerifyParams(userId,toUserId,actionType) {
		fmt.Println("params is empty..")
		return errors.New("params.isEmpty")
	}

	//change params to int
	uId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil{
		fmt.Println("uid to int failed:",err.Error())
		return
	}
	toUid, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil{
		fmt.Println("toUid to int failed:",err.Error())
		return
	}
	//package
	relation := &dto.Relation{
		UserId: uId,
		FollowId: toUid,
	}
	//choose service
	switch actionType{
	case "1":{
		r.relationDao.Action(relation)
	}
	case "2":{
		r.relationDao.CancelAction(relation)
	}
	}
	return nil
}
func (r *RelationServiceImpl)GetUserFollowList(token string,userId string)(err error,userList []vo.User){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("token verify failed:",err.Error())
		return
	}
	//check params
	if !utils.VerifyParams(userId) {
		fmt.Println("params is empty..")
		err = errors.New("params.isEmpty")
		return
	}
	uid,err := strconv.ParseInt(userId,10,64)
	if err != nil{
		fmt.Println("user id to int failed:",err.Error())
		return
	}
	//get relation
	relations := r.relationDao.GetFollowsByUserId(uid)
	//todo rpc relation to user info
	return nil,nil

}
func (r *RelationServiceImpl)GetUserFollowerList(token string,userId string)(err error,users []vo.UserVo){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		fmt.Println("token verify failed:",err.Error())
		return
	}
	//check params
	if !utils.VerifyParams(userId) {
		fmt.Println("params is empty..")
		err = errors.New("params.isEmpty")
		return
	}
	uid,err := strconv.ParseInt(userId,10,64)
	if err != nil{
		fmt.Println("user id to int failed:",err.Error())
		return
	}
	relations := r.relationDao.GetFollowersByUserId(uid)
	//todo rpc relation to user info
	return nil,nil
}
//single create
var(
	relationService RelationService
	relationServiceOnce sync.Once
)
func GetRelationService() RelationService{
	relationServiceOnce.Do(func() {
		relationService = &RelationServiceImpl{
			relationDao: dao.GetRelationDao(),
		}
	})
	return relationService
}