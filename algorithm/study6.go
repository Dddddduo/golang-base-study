package main

import "fmt"

func main() {
	var s, s1, s2, s3 int
	fmt.Scan(&s, &s1, &s2, &s3)
	if s < 425 && (s1 < 60 || s2 < 60 || s3 < 60) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
