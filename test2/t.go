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
	var arr = make([]int, n)
	for i := range n {
		fmt.Scan(&arr[i])
	}

	// 求前缀最大值
	var pre = make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = arr[i]
		if i > 0 {
			pre[i] = max(pre[i], pre[i-1])
		}
	}

	// 1.偶数位要尽可能大
	for i := 1; i < n; i += 2 {
		arr[i] = max(arr[i], pre[i])
	}

	//for i := range arr {
	//	fmt.Println(arr[i])
	//}

	var cnt int
	cnt = 0

	// 2.奇数位要尽可能小 即增加操作次数
	if arr[0] >= arr[1] {
		cnt += arr[0] - arr[1]
		cnt++
	}

	for i := 2; i < n-1; i += 2 {
		if arr[i] >= arr[i-1] || arr[i] >= arr[i+1] {
			cnt += max(arr[i]-arr[i-1], arr[i]-arr[i+1]) + 1
		}
	}

	if n%2 == 1 {
		if arr[n-1] >= arr[n-2] {
			cnt += arr[n-1] - arr[n-2]
			cnt++
		}
	}

	fmt.Println(cnt)
}
