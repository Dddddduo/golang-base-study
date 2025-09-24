package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	if n == 1 {
		fmt.Fprintln(writer, "20250121")
	} else if n == 2 {
		fmt.Fprintln(writer, "20250123")
	} else if n == 3 {
		fmt.Fprintln(writer, "20250126")
	} else if n == 4 {
		fmt.Fprintln(writer, "20250206")
	} else if n == 5 {
		fmt.Fprintln(writer, "20250208")
	} else if n == 6 {
		fmt.Fprintln(writer, "20250211")
	}
}
