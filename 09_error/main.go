package main

import (
	"errors"
	"fmt"
	strconv "strconv"
)

/*Go中的错误*/
//1.默认错误
func main01() {
	//将字符串转换为int类型，转换错误，一旦出现通过编译但是出现转换错误，错误将会被接口error捕获，且打印
	result, err := strconv.Atoi("你好")
	if err != nil {
		fmt.Println("错误原因：", err)
	} else {
		fmt.Println(result)
	}
}

//2.自定义错误信息(工厂函数)
func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("不能输入小于0的数")
	} else {
		return a + b, nil
	}
}
func main02() {
	_, err := sum(-1, -1)
	if err != nil {
		fmt.Println(err)
	}
}

//3.自定义错误信息类、实现error接口的Error
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

func main04() {
	_, err := sum2(-1, -1)
	if err != nil {
		fmt.Println(err)
	}
}

//4.error断言
func main05() {
	sum, err := sum2(-1, 2)
	if cm, ok := err.(*commonError); ok {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

}

//5.错误嵌套
