package main

import "fmt"

func main() {
	fmt.Println("老公的想法")
	var yes string
	fmt.Print("有卖西瓜的吗？(Y/N):")
	fmt.Scan(&yes)
	if yes == "Y" || yes == "y" {
		fmt.Println("买一个西瓜")
	} else {
		fmt.Println("买一个包子")
	}

	fmt.Println("----------------------------")

	fmt.Println("老婆的想法")
	fmt.Println("有卖西瓜的吗？(Y/N):")
	fmt.Scan(&yes)
	if yes == "Y" || yes == "y" {
		fmt.Println("买十个包子和一个西瓜")
	} else {
		fmt.Println("买十个包子")
	}

	print("----------------------------")
	fmt.Println("请输入考试成绩：")
	var score int
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 70 {
		fmt.Println("中等")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

}
