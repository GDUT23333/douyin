package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"sync"
	"time"
	"video_module/dao"
	"video_module/model/dto"
	"video_module/model/vo"
	"video_module/utils"
)

/**
 * @Author: Ember
 * @Date: 2022/5/10 10:50
 * @Description: TODO
 **/

type FeeService interface {
	//用户根据自己的ID获取自己的视频发布列表
	GetPublishFees(id string,token string) (list []vo.VideoVo,err error)
	//发布视频
	PublishFee(token string,id string,data *multipart.FileHeader,c *gin.Context) (count int64,err error)
}

type FeeServiceImpl struct{
	feeDao dao.FeeDao
}
func (f *FeeServiceImpl) GetPublishFees(id string,token string)(list []vo.VideoVo,err error){
	//verify token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		return
	}
	//get user fees
	fees := f.feeDao.GetFeesByUserID(userId)
	//dto change vo
	vos := make([]vo.VideoVo,len(fees))
	//get userinfo by id
	info := utils.GetUserInfo(userId)
	uservo := vo.UserVo{
		ID: info.Id,
		Name: info.Name,
		FollowCount: info.FollowCount,
		FollowerCount: info.FollowerCount,
		IsFollow: info.IsFollow,
	}
	//pacage vos
	for index,fee := range(fees){
		vos[index] = vo.VideoVo{
			ID : fee.ID,
			PlayUrl: fee.PlayUrl,
			CoverUrl: fee.CoverUrl,
			FavoriteCount: fee.ApproveCount,
			CommentCount: fee.CommentCount,
			Author: uservo,
		}
	}

	return vos,nil
}
func (f *FeeServiceImpl) PublishFee(token string,id string,data *multipart.FileHeader,c *gin.Context) (count int64,err error){
	//Verify Token
	_, _, err = utils.VerifyToken(token)
	if err != nil{
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		return
	}
	//Save Local
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName)
	err = c.SaveUploadedFile(data, saveFile)
	if err != nil{
		return
	}
	//package change dto
	video := &dto.Video{
		ID : userId,
		PlayUrl: saveFile,
		//todo  find coverurl?
		CoverUrl: "",
		ApproveCount: 0,
		CommentCount: 0,
		CreateTime: time.Now(),
	}
	return f.feeDao.CreateFee(video)
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
