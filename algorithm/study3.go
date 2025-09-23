package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	if x < 10 || x >= 100 {
		fmt.Println("No")
	} else {
		tens := x / 10
		units := x % 10

		if tens == units {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
