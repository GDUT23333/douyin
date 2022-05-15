package dto

/**
 * @Author: Ember
 * @Date: 2022/5/11 10:27
 * @Description: TODO
 **/

type VideoAction struct{
	//primary key
	ID int64
	//video key
	VideoID int64
	//user key
	UserID int64
	//action type
	ActionType int32
}