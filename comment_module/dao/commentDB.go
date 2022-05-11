package dao

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @Author: Ember
 * @Date: 2022/5/11 12:51
 * @Description: TODO
 **/

var (
	CommentDB *gorm.DB
)

//初始化数据库
func InitialDB() error{
	dsn := "root:admin@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	//生成连接池
	curDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		return errors.New("initial db failed .. ")
	}
	CommentDB = curDb
	return nil
}