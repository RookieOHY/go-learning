package main

import (
	"errors"
	"fmt"
	"golang.org/x/example/stringutil"
)

/*
1.私有的函数，使用的作用于存在于当前包
2.公有的函数，使用作用域存在于不同的包（类似于fmt）
3.一个函数一定属于一个包
*/
/*公有的为大写 私有的为小写*/
/*函数学习*/
func main() {
	fmt.Println("a")
	fmt.Println(sum(1, 1))
	fmt.Println(sum2(1, -1))
	result, err := sum2(1, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result2, err := sum3(1, -1)
	fmt.Println(result2)
	fmt.Println(sum4(1, 1))
	fmt.Println(sum5(1, 2, 4))
	fmt.Println(Sum6(1))
	fmt.Println(stringutil.Reverse("hello"))
}

/*支持定义返回值的类型*/
func sum(a int, b int) int {
	return a + b
}

/*支持定义返回的多个参数的类型*/
func sum2(a int, b int) (int, error) {
	if 0 > a || 0 > b {
		return 0, errors.New("参数不可以小于0")
	}
	return a + b, nil
}

/*支持定义返回的多个参数的名字和类型*/
func sum3(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	/*Go支持定义返回值*/
	return
}

/*支持可变参数，可以传入多个数据*/
func sum4(params ...int) int {
	sum := 0
	for _, param := range params {
		sum += param
	}
	return sum
}

/*注意点：同时存在可变参数和普通参数，要求普通参数位于可变参数之前*/
func sum5(a int, params ...int) int {
	sum := a
	for _, param := range params {
		sum += param
	}
	return sum
}
