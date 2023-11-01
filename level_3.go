package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("欢迎使用GO语言素数判断器")
	fmt.Println("请输入一个正整数")
	fmt.Scanln(&a)
	if(a==1){
		fmt.Println("不是素数")
	}
	if(a==2){
		fmt.Println("是素数")
	}
LOOP:
	// fmt.Println("123")
	for i := 2; i <= a-1; i++ {
		if a%i == 0 {
			fmt.Println("不是素数")
			break LOOP
		} else {
			fmt.Println("是素数")
			break LOOP
		}
	}
}
