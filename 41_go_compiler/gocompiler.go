package _1_go_compiler

import (
	"fmt"
	"go/scanner"
	"go/token"
)

/*
	go编译器学习
*/
// 模拟Go编译器对文本内容的扫描（内容的转换）
func ScannerText() {
	src := []byte("cos(x) + 2i*sin(x) // Euler")
	var s scanner.Scanner
	fileSet := token.NewFileSet()
	file := fileSet.AddFile("", fileSet.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)
	//执行扫描
	for {
		position, tk, lit := s.Scan()
		//EOF --> end of file 标志文件的结尾
		if tk == token.EOF {
			break
		}
		//没有到文件结尾 输出即可(位置、符号、字面量)
		fmt.Printf("%s\t%s\t%q\n", fileSet.Position(position), tk, lit)
	}
}
