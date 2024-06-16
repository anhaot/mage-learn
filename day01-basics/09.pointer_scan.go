package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入名字:")
	fmt.Scan(&name)
	fmt.Println("你输入的内容是", name)

	var age int
	fmt.Print("请输入年龄:")
	fmt.Scan(&age)
	fmt.Println("你输入的年龄是", age)
	// 如果age输入的不是数字，那么直接跳过，不会执行下面的代码
	var height float64
	fmt.Print("请输入身高:")
	fmt.Scan(&height)
	fmt.Println("你输入的身高是", height)
}
