package _5_generics

import "fmt"

/*
	Go 泛型特性学习
		1.入参是接口类型的数组 []interface{}，此时 DoRange(strings) 编译报错，原因：Cannot use 'strings' (type []string) as the type []interface{}
		2.入参是接口，此时是不报错，但是需要对入参sl进行断言后遍历
	断言：
		所谓断言，可以类比 java instance of
		x.(T):T是否实现了x接口；是，就把x接口转为T
	不足：
		断言：转具体类型，当入参是 []int64 类型时，编译器不会报错，当时运行时会报错，原因：interface conversion: interface {} is []int64, not []string
	泛型：
		调用方来设置类型，而不是被调用方限制死类型
		语法：
			例子：func methodName[T any 或者 string | int64](param []T){}
			格式：[约束的类型]
			内置的约束类型：
				- any : 允许实参是go里面的所有类型 等于interface{}
				- comparable :允许实参是内置的可以比较的类型： int uint float bool struct 指针等
	泛型类型：


*/

// 类型定义
type cInt []int
type cStr []string

type Slice[T int | string] []T

type cMap[KEY comparable, VALUE comparable] map[KEY]VALUE

func Init() {
	var il cInt = []int{25, 26}
	var sl cStr = []string{"subhee", "rookie"}
	var itl Slice[int] = []int{25, 26}
	var stl Slice[string] = []string{"subhee", "rookie"}

	var cmap cMap[int64, int64] = map[int64]int64{
		23: 23,
	}

	fmt.Println(len(il))
	fmt.Println(len(sl))
	fmt.Println(len(itl))
	fmt.Println(len(stl))

	for k, v := range cmap {
		fmt.Println(k, v)
	}

	// 打印 itl 和 stl 的类型
	fmt.Printf("%T", itl) // Slice[int]
	fmt.Printf("%T", stl) // Slice[string]
	fmt.Printf("%T", cmap)

}

// DoRangeWith 带泛型
func DoRangeWith[T any](sl []T) {
	for _, v := range sl {
		println(v)
	}
}

func DoRange(sl interface{}) {

	//for _, v := range sl {
	//	//println(i)
	//	println(v)
	//}

	for _, v := range sl.([]string) {
		//println(i)
		println(v)
	}

}

// Range 遍历数组
func Range() {
	names := []string{"rk", "subhee"}
	ages := []int64{26, 27}
	//DoRange(names)
	//DoRange(ages)
	DoRangeWith(names)
	DoRangeWith(ages)
}
