package main

import "fmt"

func main() {
	// 从0开始
	const (
		E1 int = iota
		E2
		E3
	)
	// 不使用0
	const (
		E4 int = 3
		E5     = iota
		E6
	)
	fmt.Println(E1, E2, E3, E4, E5, E6)
}

//结果：0 1 2 3 1 2
