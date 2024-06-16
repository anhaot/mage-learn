package main

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto label
		}
		println(i)
	}
label: // 大写，可以随意写
	println("label")
}
