package main

import (
	"bufio"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	lineT, _ := reader.ReadString('\n')
	lineT = strings.TrimSpace(lineT)
	t, _ := strconv.Atoi(lineT)
	for i := 0; i < t; i++ {
		dduo()
	}
}

func dduo() {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	ans := 0
	j := 0
	for i := 0; i < n; i++ {
		if i >= j && s[i] == '1' {
			ans++
		}
		if s[i] == '1' {
			if i+k > j {
				j = i + k
			}
		}
		//writer.WriteString(strconv.Itoa(ans))
	}
	writer.WriteString(strconv.Itoa(ans) + "\n")
}
