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
	var n int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	a := make([]int, n)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(parts[i])
	}

	sort.Ints(a)

	var ans, t int64
	t = int64(a[0])

	for i := 1; i < n; i++ {
		ans += t * int64(a[i])
		t += int64(a[i])
	}

	fmt.Println(ans)
}
