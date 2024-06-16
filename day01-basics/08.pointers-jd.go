package main

import "fmt"

func main() {
	A := 2
	B := A
	B = 3
	fmt.Println(A, B)

	// 声明了一个名为cc的指针变量，并使用&A（即A的地址）为其赋值
	var cc *int = &A
	// 声明了一个名为C的指针变量，并同样使用&A为其赋值。
	C := &A
	fmt.Printf("%T %T %p\n", C, cc, cc) // *int *int 0xc000126058
	// cc来访问cc所指向的值，即A的值。因为A的值为2，所以输出为2
	fmt.Println(*cc)
	// 修改了cc所指向的值（即A的值）为4。然后，我们打印A的值，输出为4
	*cc = 4
	fmt.Println(A)
}
