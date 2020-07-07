package util

import (
	"encoding/json"
	"log"
	"net/http"
)

//定义一个结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}


//失败的返回结果
func Fail(writer http.ResponseWriter, msg string) {
	response(writer, 512, nil, msg)
}

//返回成功
func OK(writer http.ResponseWriter, data interface{}, msg string) {
	response(writer, 200, data, msg)
}

func response(writer http.ResponseWriter, code int, data interface{}, msg string) {
	//设置header 为JSON 默认是test/html,所以特别指出返回的数据类型为application/json
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	rep := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//将结构体转化为json字符串
	ret, err := json.Marshal(rep)
	if err != nil {
		log.Panicln(err.Error())
	}

	//返回json ok
	writer.Write(ret)
}


