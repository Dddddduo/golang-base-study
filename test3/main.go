package main

import (
	"fmt"
)

var (
	mod = 1000000007
	//mod = 998244353
)

func main() {
	var t int
	t = 1
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		dduo()
	}
}

func dduo() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, 100)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x < 100 {
			arr[x]++
		}
	}
	ans := 0
	for m := 0; m < 100; m++ {
		if arr[m] == 0 {
			ans = m
			break
		}
	}
	fmt.Println(ans)
}
