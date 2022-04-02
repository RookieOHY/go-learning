package main

import (
	"fmt"
)

/*匿名函数和闭包*/
func main01() {
	/*
		1.sum不是函数而是变量
		2.需要声明返回值类型（否则会报错）
	*/
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(1, 1))
}
func main02() {
	result := out()
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
}

/*闭包概念：在函数嵌套中，内部函数（匿名函数）可以使用外部函数的变量*/
func out() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/* GO方法必须要有一个接收者，接收者是类型，这样方法就和这个类型绑定在一起，称为这个类型的方法*/
type Age uint

/*新定义了一个类型为Age,Age类型的方法为String()*/
/*方法的调用：接受一个本地变量（称之为接受者），使用“.”调用这个变量的方法*/
func (age Age) String() {
	fmt.Println("the age is", age)
}
func (age *Age) Modify() {
	*age = Age(30)
}

func main03() {
	age := Age(25)
	age.String() //25
	age.Modify() //30
	age.String() //30
	(&age).Modify()
	(&age).String() //30

}

/*方法表达式*/
func main() {
	age := Age(25)
	sm := Age.String
	sm(age)
}
