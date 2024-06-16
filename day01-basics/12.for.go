package main

import "fmt"

func main() {
	sum := 0
	// 初始化子语句；条件子语句；后置子语句；
	//1.初始化-->2.条件子语句-->3.后置语句->4.循环体-->5.条件子语句-->6.后置语句-->7.循环体...
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("--------------------")

	sum = 0
	i := 1
	for i <= 100 {
		i++
	}
	fmt.Println(sum)

	fmt.Println("-------死循环---------")

	/*	i = 0
		for {
			fmt.Println(i)
			i++
		}*/

	fmt.Println("--------for-range---------")

	// 字符串、切片、数组、map遍历
	strings := "我爱祖国"
	for i, ch := range strings {
		fmt.Printf("%T %d %T ,%q\n", i, i, ch, ch)
	}
}
