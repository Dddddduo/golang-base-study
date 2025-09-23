package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x+y > 9 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
