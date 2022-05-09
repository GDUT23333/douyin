package model

/**
 * @Author: Ember
 * @Date: 2022/5/9 15:05
 * @Description: TODO
 **/

type Response struct{
	StatusCode int64 `json:"status_code"`
	StatusMsg string `json:"status_msg,omitempty"`
}

