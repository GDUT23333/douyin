package service

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"
	"video_module/dao"
	"video_module/model/dto"
	"video_module/utils"
	"video_module/video_rpc_service"
)

/**
 * @Author: Ember
 * @Date: 2022/5/10 10:50
 * @Description: TODO
 **/

type FeeService interface {
	//推送视频
	PushFeed(latestTime int64)(list []*video_rpc_service.Video,nextTime int64,err error)
	//用户根据自己的ID获取自己的视频发布列表
	GetPublishFees(id int64,token string) (list []*video_rpc_service.Video,err error)
	//发布视频
	PublishFee(token string,id int64,data []byte) (count int64,err error)
}

type FeeServiceImpl struct{
	feeDao dao.FeeDao
}
func (f *FeeServiceImpl) GetPublishFees(id int64,token string)(list []*video_rpc_service.Video,err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		return
	}
	//get user fees
	fees := f.feeDao.GetFeesByUserID(id)
	//dto change vo
	vos := make([]*video_rpc_service.Video,len(fees))
	// TODO  RPC Service
	//get userinfo by i
	//pacage vos
	for index,fee := range(fees){
		vos[index] = &video_rpc_service.Video{
			Id : &fee.ID,
			PlayUrl: &fee.PlayUrl,
			CoverUrl: &fee.CoverUrl,
			FavoriteCount: &fee.ApproveCount,
			CommentCount: &fee.CommentCount,
			Author: nil,
		}
	}

	return vos,nil
}
func (f *FeeServiceImpl) PublishFee(token string,id int64,data []byte) (count int64,err error){
	//Verify Token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		return
	}
	//Save Local
	filename := filepath.Base("test")
	finalName := fmt.Sprintf("%d_%s", id, filename)
	saveFile := filepath.Join("./public/", finalName)
	//permissions := 0644
	//err := ioutil.WriteFile(saveFile, data, permissions)
	if err != nil{
		return
	}
	//package change dto
	video := &dto.Video{
		ID : id,
		PlayUrl: saveFile,
		//todo  where get coverurl?
		CoverUrl: "",
		ApproveCount: 0,
		CommentCount: 0,
		CreateTime: time.Now(),
	}
	return f.feeDao.CreateFee(video)
}
func(f *FeeServiceImpl )PushFeed(latestTime int64)(list []*video_rpc_service.Video,nextTime int64,err error){
	return nil,0,nil
}
var(
	feeService FeeService
	feeServiceOnce sync.Once
)
// single create
func GetFeeService() FeeService{
	feeServiceOnce.Do(func() {
		feeService = &FeeServiceImpl{
			feeDao: dao.GetFeeDao(),
		}
	})
	return feeService
}
