package main

import "fmt"

func main() {
	// 作用域：定义标识符可以使用的范围
	// 在go中用{}来定义作用域的范围
	scopes := "这是作用域"
	{
		zuoyong := 2
		fmt.Println(scopes, zuoyong)
	}
}

/*
在大括号外部访问不到scopes
fmt.Println(scopes)
*/
