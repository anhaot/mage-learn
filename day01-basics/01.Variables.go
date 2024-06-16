package main

import "fmt"

// 全局变量定义
var C9 = "hello C9"
var C10 string = "hello C10"
var (
	C11 string  = "hello C11"
	C12 int     = 23
	C13 float32 = 3.14
)

/*
在全局变量中，这个不可以这么使用
C11 := "hello C11"
*/

func main() {
	// 简单局部变量定义
	C1 := 2
	var C2 = "Hello C2"
	var C3 string = "Hello C3"
	// 相同类型局部变量定义
	var C4, C5 string = "Hello C4", "Hello C5"
	// 多个局部变量定义
	var (
		C6 string = "Hello C6"
		C7 int
		C8 = 3.15
	)

	fmt.Println(C1, C2, C3, C4, C5, C6, C7, C8, C9, C10, C11, C12, C13)
}

// 结果：2 Hello C2 Hello C3 Hello C4 Hello C5 Hello C6 0 3.15 hello C9 hello C10 hello C11 23 3.14
