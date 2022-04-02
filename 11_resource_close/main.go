package main

import (
	"os"
)

//释放资源: Go在任何情况（异常或者正常运行）下都应该释放资源
//defer 用于修饰一个函数或者是方法，保证函数在return前执行 file.Close()方法
func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() //释放文件资源
	return nil, err
}
func main() {

}
