package simple

import (
	"errors"
	"fmt"
	"log"
)

/*
	Go错误处理方式的最佳实践（简单版）
		利用error包定义一些常用的错误信息，在业务方法出现错误的时候，获取对应的error,判断error类型，以输出不同的错误信息。
*/

/*步骤*/
/*1.定义常用的错误信息*/
var (
	errEntryNotFound = errors.New("entry not found")
	errMissingParam  = errors.New("missing param")
	errUnknown       = "errUnknown"
)

func NewMap() map[string]string {
	initMap := make(map[string]string)
	initMap["name"] = "RookieOHY"
	initMap["age"] = "25"
	return initMap
}
func getItem(key string) (string, error) {
	newMap := NewMap()
	val, ok := newMap[key]
	if !ok {
		//未找到entry 错误
		return "", errEntryNotFound
	}
	return val, nil
}
func ServiceName(key string) {
	val, err := getItem(key)
	if err != nil {
		//调用处理方法，判断类型
		handleError(err)
		return
	}
	//正常的业务
	fmt.Println("我是正常的业务！我的执行结果为-->", val)
	return
}

//定义一个全局的错误处理方法
func handleError(err error) {
	//判断、执行其他业务（日志等）结束
	if err == nil {
		return
	}
	switch err {
	case errMissingParam:
		log.Println("我是错误的业务！请求参数缺失")
		fmt.Println(errMissingParam.Error())
	case errEntryNotFound:
		log.Println("我是错误的业务！未查询到数据")
		fmt.Println(errEntryNotFound.Error())
	default:
		log.Println("我是错误的业务！出现未知错误")
		fmt.Println(errUnknown)
	}
}
