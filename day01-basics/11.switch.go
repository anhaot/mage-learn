package main

import "fmt"

func main() {
	fmt.Println("老公的想法")
	var yes string
	fmt.Print("有卖西瓜的吗？(Y/N):")
	fmt.Scan(&yes)
	switch yes {
	case "Y", "y":
		fmt.Println("买一个西瓜")
	default:
		fmt.Println("买一个包子")
	}
	fmt.Println("-----------------------")
	fmt.Println("老婆的想法")
	var yess string
	fmt.Print("有卖西瓜的吗？(Y/N):")
	fmt.Scan(&yess)
	switch yess {
	case "Y", "y":
		fmt.Println("买一个西瓜和十个包子")
	default:
		fmt.Println("买十个个包子")
	}

	fmt.Println("-----------------------")
	fmt.Println("请输入考试成绩：")
	var score int
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 70:
		fmt.Println("中等")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
