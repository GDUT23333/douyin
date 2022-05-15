package utils

import (
	"github.com/dgrijalva/jwt-go"
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

