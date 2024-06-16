package main

import "fmt"

func main() {
	num := 10
	num01 := 3.14
	num02 := "hello"
	fmt.Printf("%T %d \n", num, num)
	// %f 来打印浮点数
	fmt.Printf("%T %f \n", num01, num01)
	// %s 来打印字符串
	fmt.Printf("%T %s \n", num02, num02)
	// number 函数的返回值
	c := number() // number 函数以返回 int 类型的值
	fmt.Printf("%T %d \n", c, c)

	// strings 函数的返回值
	var d string = strings()
	fmt.Printf("%T %s \n", d, d)
}

// number 函数以返回 int 类型的值
func number() int {
	a, b := 1, 2
	c := a + b
	return c
}

func strings() string {
	a := "hello"
	b := "world"
	c := a + b
	return c
}
