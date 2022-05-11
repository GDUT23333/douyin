package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
	"strings"
	"time"
	"user_module/model/dto"
)

/**
 * @Author: Ember
 * @Date: 2022/5/9 20:28
 * @Description: TODO
 **/

//校验字符串参数
func VerifyParams(strs ...string) bool{
	for _,str := range strs{
		if !verify(str){
			return false
		}
	}
	return true
}
func verify(str string) bool{
	if str == ""{
		return false
	}
	return true
}
//md5加密
func Md5Encryption(str string) (result string,err error){
	m := md5.New()
	_,err = io.WriteString(m,str)
	if err != nil {
		log.Fatal(err)
		return
	}
	arr := m.Sum(nil)
	// 将编码转换为字符串
	newArr := fmt.Sprintf("%x",arr)
	//输出字符串字母都是小写，转换为大写
	return strings.ToTitle(newArr),nil
}


var(
	//jwt salt value
	jwtkey = []byte("douyin")
)

//verify token
func VerifyToken(token string)(*jwt.Token, *jwt.StandardClaims, error){
	Claims := &jwt.StandardClaims{}
	t, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return t, Claims, err
}
