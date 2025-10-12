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

	//scanner.Scan()
	//t, _ := strconv.Atoi(scanner.Text())
	
	t := 1

	for i := 0; i < t; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		n, _ := strconv.Atoi(parts[0])
		k, _ := strconv.Atoi(parts[1])
		if n-k < 1 {
			fmt.Fprintln(writer, "NO")
		} else {
			fmt.Fprintln(writer, "YES")
		}
	}
}
