package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @Author: Ember
 * @Date: 2022/5/8 22:03
 * @Description: TODO
 **/
type UserInfo struct{
	ID int `grom:"column:id"`
	UserName string `grom:"column:user_name"`
	UserPassWord string `grom:"column:user_password"`
	UserNick string `grom:"user_nick"`
}
//实现Table接口来更改默认表名
//这个表名不会进行动态变化，而是会缓存下来
func (u *UserInfo) TableName() string{
	return "user_info"
}

func main(){
	dsn := "root:admin@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		return
	}
	var user UserInfo
	db.First(&user)
	fmt.Println("user:%v",user)

}



