package main

import "fmt"

func longestPalindrome(s string) string {
	var res string

	// 奇数
	for i := 0; i < len(s); i++ {
		var subs string
		for j := 0; j <= min(i, len(s)-i-1); j++ {
			if s[i-j] != s[i+j] {
				break
			}
			subs = s[i-j : i+j+1]
			if len(subs) > len(res) {
				res = subs
			}
		}
	}

	// 偶数
	for i := 1; i < len(s); i++ {
		var subs string
		for j := 0; j < min(i, len(s)-i); j++ {
			if s[i-1-j] != s[i+j] {
				break
			}
			subs = s[i-1-j : i+j+1]
			if len(subs) > len(res) {
				res = subs
			}
		}
	}

	return res
}

func main() {
	fmt.Println(longestPalindrome("babab"))
	fmt.Println("====")
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println("====")
	fmt.Println(longestPalindrome("a"))
	fmt.Println("====")
	fmt.Println(longestPalindrome("aa"))
	fmt.Println("====")
	fmt.Println(longestPalindrome("ac"))
	fmt.Println("====")
	fmt.Println(longestPalindrome("aaa"))
	fmt.Println("====")
	fmt.Println(longestPalindrome("aacabdkacaa"))
}
