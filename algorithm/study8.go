package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max(max(a, b), c)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	//scanner.Scan()
	//x, _ := strconv.Atoi(scanner.Text())
	//scanner.Scan()
	//y, _ := strconv.Atoi(scanner.Text())
	//scanner.Scan()
	//z, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, " ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])

	sum := x + y + z

	m := max3(x, y, z)

	fmt.Fprintln(writer, max(sum-m, m))
}
