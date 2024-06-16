package main

import "fmt"

// 全局常量
const C7 = 7
const C8 string = "Hello C8"
const Co9, Co10 int = 9, 10
const (
	Co11        = 11
	Co12 string = "Hello C12"
)
const (
	Co16 int = 16
	Co17
	Co18
)

func main() {
	const C1 = 1
	const C2 string = "Hello C2"
	// 相同类型多个常量
	const C3, C4 int = 3, 4
	// 不同类型多个常量
	const (
		C5        = 5
		C6 string = "Hello C6"
	)
	// 省略常量
	const (
		Co13 int = 11
		Co14
		co15
	)
	fmt.Println(C1, C2, C3, C4, C5, C6, C7, C8, Co9, Co10, Co11, Co12, Co13, Co14, co15, Co16, Co17, Co18)
}

// 结果：1 Hello C2 3 4 5 Hello C6 7 Hello C8 9 10 11 Hello C12 11 12 13 14 15 16 17 18
