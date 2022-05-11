package dao

import (
	"relation_module/model/dto"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 15:25
 * @Description: TODO
 **/
type RelationDao interface{
	//get user follow
	GetFollowsByUserId(userId int64) []*dto.Relation
	//get user follower
	GetFollowersByUserId(userId int64) []*dto.Relation
	// action
	Action(relation *dto.Relation)
	//cancel Action
	CancelAction(relation *dto.Relation)
}
type RelationDaoImpl struct{

}

func (r *RelationDaoImpl) GetFollowsByUserId(userId int64) []*dto.Relation {

}

func (r *RelationDaoImpl) GetFollowersByUserId(userId int64) []*dto.Relation {

}

func (r *RelationDaoImpl)Action(relation *dto.Relation){

}

func (r *RelationDaoImpl)CancelAction(relation *dto.Relation){

}

//single create
var(
	relationDao RelationDao
	relaationDaoOnce sync.Once
)
func GetRelationDao() RelationDao{
	relaationDaoOnce.Do(func() {
		relationDao = &RelationDaoImpl{}
	})
	return relationDao
}
