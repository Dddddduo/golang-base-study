package main

import (
	"fmt"
)

const mod = 998244353

func main() {
	var n int
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println(0)
		return
	}

	result := 1
	for i := 3; i <= n+1; i++ {
		result = (result * i) % mod
	}

	fmt.Println(result)
}
