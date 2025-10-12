package main

import (
	"fmt"
)

func main() {
	scores := [7]int{0, 0, 0, 0, 0, 0, 0}
	eliminated := [7]bool{false, false, false, false, false, false, false}

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var op, x, y int
		fmt.Scan(&op, &x, &y)

		if eliminated[x] {
			continue
		}

		if op == 1 {
			scores[x] += y
		} else if op == 2 {
			scores[x] += y
			eliminated[x] = true
		}
	}

	successCount := 0
	for i := 1; i <= 6; i++ {
		if !eliminated[i] {
			successCount++
		}
	}

	bonus := 100 * successCount
	for i := 1; i <= 6; i++ {
		if !eliminated[i] {
			scores[i] += bonus
		}
	}

	maxScore := -1
	champion := 0
	for i := 1; i <= 6; i++ {
		if scores[i] > maxScore {
			maxScore = scores[i]
			champion = i
		}
	}

	fmt.Println(champion)
}
