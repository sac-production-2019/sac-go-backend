package utils

import(
	"github.com/sac-production-2019/sac-go-backend/model"
)

type Result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type page struct{
	Count int `json:"count"`
	Items interface{} `json:"items"`
}

const (
	success int =200
	fail int =400
	unAuth int = 330 //无权限
	errJWT int = 340 //JWT 未通过验证
	erorr int = 500
)

type LoginResult struct {
	Token string `json:"token"`
	model.User
}

func newRes(code int, msg string , data ...interface{}) (int, Result){
	if len(data)>0{
		return 200,Result{
			Code: code,
			Msg: msg,
			Data: data[0],
		}
	}
	return 200,Result{
		Code: code,
		Msg:msg,
	}
}


func Success(msg string, data ... interface{})(int,Result){
	return newRes(success,msg,data...)
}

func Fail(msg string, data...interface{})(int, Result){
	return newRes(fail,msg,data...)
}

func UnAuth(msg string,data...interface{}) (int,Result){
	return newRes(unAuth,msg,data...)
}

func ErrJWT(msg string,data...interface{})(int,Result){
	return newRes(errJWT,msg,data...)
}

func Error(msg string, data...interface{}) (int, Result){
	return newRes(erorr,msg,data...)
}

func NewPage(msg string, items interface{},count int) (int,Result){
	return 200,Result{
		Code: success,
		Msg: msg,
		Data: page{
			Items:items,
			Count: count,
		},
	}
}

