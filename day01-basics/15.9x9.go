package main

import "fmt"

func main() {
	for num := 1; num <= 9; num++ {
		for i := 1; i <= num; i++ {
			fmt.Printf("%d * %d = %2d\t", i, num, i*num)
		}
		fmt.Println()
	}
}
