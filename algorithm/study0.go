package main

import "fmt"

func main() {
	var dduo int
	fmt.Scan(&dduo)

	if dduo < 3 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(1, dduo-1)
	}
}
