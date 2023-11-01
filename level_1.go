package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("欢迎使用GO语言加法器")
	fmt.Println("请输入两个整数")
	fmt.Println("请输入第一个整数")
	fmt.Scanln(&a)
	fmt.Println("请输入第二个整数")
	fmt.Scanln(&b)
	fmt.Println("计算结果如下：")
	fmt.Println(a + b)
}
