package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	votes := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 6; i++ {
		scanner.Scan()
		food := scanner.Text()
		votes[food]++
	}

	maxVotes := 0
	for _, count := range votes {
		if count > maxVotes {
			maxVotes = count
		}
	}

	for food, count := range votes {
		if count == maxVotes {
			fmt.Println(food)
		}
	}
}
