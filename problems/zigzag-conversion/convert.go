package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	// ジグザグに渡される文字を格納するスライス
	container := make([][]string, numRows)

	// コンテナを1次元均したスライス
	var flatten []string

	// 何番目の文字が何列目に格納されるか
	positions := make(map[int]int)

	// 格納先が変わる周期
	pattern := max(numRows*2-2, numRows)

	// positions と pattern を対応付け
	for i := 0; i < pattern; i++ {
		positions[i] = min(pattern-i, i)
	}

	// 渡された文字列をコンテナに格納していく
	for i, v := range s {
		container[positions[i%pattern]] = append(container[positions[i%pattern]], string(v))
	}

	// flatten に、スライスごとに連結した文字列を格納
	for _, v := range container {
		flatten = append(flatten, strings.Join(v, ""))
	}

	// flatten を連結した文字列を返却
	return strings.Join(flatten, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
}
