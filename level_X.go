package main

import (
	"fmt"
	"math/rand"
)

func main(){
	i:=rand.Intn(200)
	var ans [200] int
	for n:= 1 ; n<200 ; n++{
		ans[n]=n
	}
	var e int;
	e=guess(ans,i)
	fmt.Println("随机数是：")
	fmt.Println(e)
}
func guess(a [200]int ,b int) int{
	start :=0
	end := len(a)-1
	for start <=end{
		mid:=(start + end)/2
		if a[mid] < b {
			start = mid +1
		} else if a[mid] >b {
			end =mid -1
		} else {
			return mid
		}
	}
	return 0
}