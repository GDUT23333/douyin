package dao

import (
	"sync"
	"video_module/model/dto"
)

/**
 * @Author: Ember
 * @Date: 2022/5/10 10:50
 * @Description: TODO
 **/

type FeeDao interface{
	//用户根据自己的ID获取自己的视频发布列表
	GetFeesByUserID(id int64) []*dto.Video
	//发布视频
	CreateFee(video *dto.Video) (count int64,err error)
}
type FeeDaoImpl struct{

}
func (*FeeDaoImpl) GetFeesByUserID(id int64) []*dto.Video{
	slice := make([]*dto.Video,10)
	FeeDB.Where("user_id = ?",id).Find(&slice)
	return slice
}

func (*FeeDaoImpl) CreateFee(video *dto.Video) (count int64,err error){
	result := FeeDB.Create(&video)
	count = result.RowsAffected
	err = result.Error
	return
}

var(
	feeDao FeeDao
	feeDaoOnce sync.Once
)

//single create
func GetFeeDao() FeeDao{
	feeDaoOnce.Do(func() {
		feeDao = &FeeDaoImpl{}
	})
	return feeDao
}
