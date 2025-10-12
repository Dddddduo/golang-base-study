package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		solve(scanner, writer)
	}
}

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	numStrs := strings.Fields(scanner.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(numStrs[i])
	}

	sort.Ints(a)

	var s int
	for i := 2; i < n; i++ {
		s += a[i] * 2
	}
	s += a[0]
	s += a[1]

	fmt.Fprintln(writer, s)
}
