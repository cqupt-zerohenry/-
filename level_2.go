package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Println("欢迎使用GO语言圆面积计算器")
	fmt.Println("请输入圆的半径")
	fmt.Scanln(&r)
	fmt.Println("计算结果(保留两位小数)如下")
	fmt.Printf("%.2f", math.Pi*r*r)
}
