package main

import (
	"fmt"
	"math"
)

const epsilon = 1e-8 // 用于浮点数比较的精度

func main() {
	var k1, b1, k2, b2 float64

	fmt.Scan(&k1, &b1, &k2, &b2)

	if math.Abs(k1-k2) < epsilon {
		if math.Abs(b1-b2) < epsilon {
			fmt.Println("Wrong Answer")
		} else {
			fmt.Println("Ping Xing")
		}
	} else {
		x := (b2 - b1) / (k1 - k2)
		y := k1*x + b1
		fmt.Printf("%.8f %.8f\n", x, y)
	}
}
