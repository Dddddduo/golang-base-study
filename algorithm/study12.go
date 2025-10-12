package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	x := make([]int64, n)
	scanner.Scan()
	xStrs := strings.Fields(scanner.Text())
	for i := 0; i < n; i++ {
		val, _ := strconv.ParseInt(xStrs[i], 10, 64)
		x[i] = val
	}

	y := make([]int64, n)
	for i := 0; i < n; i++ {
		y[i] = 0
		for j := i + 1; j < n; j++ {
			if x[j] > x[i] {
				y[i] = int64(j + 1)
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", y[i])
	}
}
