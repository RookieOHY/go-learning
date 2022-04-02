package main

import (
	"errors"
	"fmt"
	"strconv"
)

//错误嵌套
//使用场景：调用一个函数，返回了一个错误信息 error，在不想丢失这个 error 的情况下，又想添加一些额外信息返回新的 error
//使用示例：
//1.需求转换为结构体
type MyError struct {
	err error
	msg string
}

//2.实现error接口
func (e *MyError) Error() string {
	return e.err.Error() + e.msg
}

func main01() {
	_, err := strconv.Atoi("你好")
	myError := MyError{
		err: err,
		msg: "转换失败",
	}
	fmt.Println(myError.err, myError.msg)
}

//2.go 1.13版本提供的错误嵌套(ErrorWrapping)
func main02() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
	//解开错误
	fmt.Println(errors.Unwrap(w))
}

//3.判断是否为同一个error
func main03() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
	fmt.Println(errors.Is(w, e))
}

//4.有了error嵌套之后，原本的判断两个error是不是同一个error的方法以及error的断言将
//5.使用errors.As实现断言
type commonError struct {
	errorCode int    //错误码
	errorMsg  string //错误信息
}

func (ce *commonError) Error() string {
	return ce.errorMsg
}
func sum2(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{
			errorCode: 500,
			errorMsg:  "服务器错误",
		}
	} else {
		return a + b, nil
	}
}
func main() {
	var cm *commonError
	i, err := sum2(-1, -1)
	if errors.As(err, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(i)
	}

}

//6.
