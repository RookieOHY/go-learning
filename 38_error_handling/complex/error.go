package complex

import (
	"fmt"
	"go-learning/38_error_handling/simple"
)

/*
	Go错误处理方式的最佳实践（自定义版）
		自定义error结构体 判断错误类型 返回对应的自定义结构体对象
*/
//定义新类型
type errType string

//定义类型为新的错误类型的常量
const (
	errNotFound     errType = "item not found"
	errUnKnown      errType = "unknown error"
	errMissingParam errType = "missing param"
)

// BusinessError 自定义错误结构体
type BusinessError struct {
	errType errType
	msg     string
}

func (be *BusinessError) Error() string {
	return fmt.Sprintf("%v->%v", be.errType, be.msg)
}

func NewBusinessError(errType errType, msg string) *BusinessError {
	return &BusinessError{errType: errType, msg: msg}
}

func getItem(key string) (string, error) {
	cn := simple.NewMap()
	value, ok := cn[key]
	if !ok {
		//无对应的值
		return "", NewBusinessError(errNotFound, key)
	}
	return value, nil
}

//自定义错误的判断
func switchErrorTypeChecking(err error) {
	//存在错误 那么判断入参的类型
	if err != nil {
		switch err.(type) {
		//是自定义类型
		case *BusinessError:
			//打印实现Error接口的方法
			fmt.Println(err.Error())
		//不是自定义类型
		default:
			fmt.Println(errUnKnown)
		}
	}
}

// HandleErrorWithCustom 业务方法
func HandleErrorWithCustom(key string) {
	item, err := getItem(key)
	if err != nil {
		//传入错误，判断错误是自定义的还是非自定
		switchErrorTypeChecking(err)
		return
	}
	fmt.Println("业务正常..", item)
	return
}
