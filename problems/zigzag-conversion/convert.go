package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	var container [][]string
	var flatten []string
	pattern := max(numRows*2-2, numRows)
	positions := make(map[int]int)

	for i := 0; i < pattern; i++ {
		positions[i] = min(pattern-i, i)
		container = append(container, nil)
	}

	for i, v := range s {
		container[positions[i%pattern]] = append(container[positions[i%pattern]], string(v))
	}

	for _, v := range container {
		flatten = append(flatten, strings.Join(v, ""))
	}

	return strings.Join(flatten, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
}
