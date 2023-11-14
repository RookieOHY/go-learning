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
		声明泛型类型，可以避免定义很多种类型。
		语法：
			type typeName[T 类型1 | 类型2] []T
		注意:
			假设定义了一种泛型类型：type Slice[T int | string] []T
			使用的时候：Slice[int] 和 Slice[string] 对应的变量 是2种 不同的变量。无法做比较。
			声明变量的时候，指定的类型不在约定的类型里面，发生错误
	泛型结构体：
		语法：type structName[T 类型1 | 类型2] struct{}
	泛型接口：
		语法：type interfaceName[T string | int] interface {}
	泛型函数：
		可以使用的地方：
			1.泛型reciever
			2.泛型函数
			3.泛型类型
	自定义泛型类型：


*/

// MyInterface 泛型
type MyInterface[T string | int] interface {
	Print(data T)
}

// MyInterface 接口的具体类型
type StringPrinter struct{}

func (sp StringPrinter) Print(data string) {
	fmt.Println("String data:", data)
}

// MyInterface 接口的具体类型
type IntPrinter struct{}

func (ip IntPrinter) Print(data int) {
	fmt.Println("Integer data:", data)
}

func InitInterface() {
	// 使用 StringPrinter 实例化 MyInterface
	var strPrinter MyInterface[string] = StringPrinter{}
	strPrinter.Print("Hello")

	// 使用 IntPrinter 实例化 MyInterface
	var intPrinter MyInterface[int] = IntPrinter{}
	intPrinter.Print(42)
}

// MyStruct 泛型结构体
type MyStruct[T string | int] struct {
	Id   T
	Name string
}

func InitMyStruct() {
	// 使用字符串类型
	strStruct := MyStruct[string]{Id: "abc", Name: "John"}

	// 使用整数类型
	intStruct := MyStruct[int]{Id: 123, Name: "Doe"}

	// 打印结构体内容
	fmt.Printf("String Struct: %+v\n", strStruct)
	fmt.Printf("Int Struct: %+v\n", intStruct)
}

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

// 泛型 receiver
func (slc Slice[T]) Sum() T {
	var sum T
	for _, v := range slc {
		sum += v
	}
	return sum
}

func DoSum() {
	var slc Slice[int] = []int{1, 2}
	println(slc.Sum())
}

// 泛型函数
func Add(a int, b int) int {
	return a + b
}

func AddT[T int | float64 | string](a T, b T) T {
	return a + b
}

func DoAddT() {
	//使用方式：直接声明形参是什么类型的 或者 不指定形参（自我推断类型）
	fmt.Println(AddT(1, 2))
	fmt.Println(AddT[int](1, 2))

	fmt.Println(AddT("1", "2"))
	fmt.Println(AddT[string]("1", "2"))
}

// 自定义泛型类型 : 当函数的入参可以有很多形参的类型时，可能要写很长的申明
type MyType interface {
	int | ~int8 | int16 | int32 | int64
}

func GetMaxNum[T MyType](a, b T) T {
	if a > b {
		return a
	}
	return b
}
