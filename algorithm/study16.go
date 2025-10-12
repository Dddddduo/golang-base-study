package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	count := 0
	temp := n
	for temp > 0 {
		if temp%2 == 1 {
			count++
		}
		temp /= 2
	}

	maxCapacity := (1 << k) - 1

	// 判断是否能装下
	if count <= k && n <= maxCapacity {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
