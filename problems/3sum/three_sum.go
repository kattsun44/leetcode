package main

import (
	"fmt"
	"slices"
)

type keys struct {
	key1 int
	key2 int
	key3 int
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)
	maps := make(map[keys]bool)
	combination := combination(nums, 3)
outer:
	for _, comb := range combination {
		if maps[keys{comb[0], comb[1], comb[2]}] {
			continue outer
		}
		maps[keys{comb[0], comb[1], comb[2]}] = true
		sum := 0
		for _, v := range comb {
			sum += v
		}
		if sum == 0 {
			res = append(res, comb)
		}
	}
	return res
}

// ref: https://obelisk.hatenablog.com/entry/2018/10/09/025240
func combination(ar []int, n int) (result [][]int) {
	if n <= 0 || len(ar) < n {
		return
	}
	if n == 1 {
		for _, a := range ar {
			result = append(result, []int{a})
		}
	} else if len(ar) == n {
		result = append(result, ar)
	} else {
		for _, a := range combination(ar[1:], n-1) {
			result = append(result, append([]int{ar[0]}, a...))
		}
		result = append(result, combination(ar[1:], n)...)
	}
	return
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(threeSum([]int{-13, 5, 13, 12, -2, -11, -1, 12, -3, 0, -3, -7, -7, -5, -3, -15, -2, 14, 14, 13, 6, -11, -11, 5, -15, -14, 5, -5, -2, 0, 3, -8, -10, -7, 11, -5, -10, -5, -7, -6, 2, 5, 3, 2, 7, 7, 3, -10, -2, 2, -12, -11, -1, 14, 10, -9, -15, -8, -7, -9, 7, 3, -2, 5, 11, -13, -15, 8, -3, -7, -12, 7, 5, -2, -6, -3, -10, 4, 2, -5, 14, -3, -1, -10, -3, -14, -4, -3, -7, -4, 3, 8, 14, 9, -2, 10, 11, -10, -4, -15, -9, -1, -1, 3, 4, 1, 8, 1}))
	fmt.Println(threeSum([]int{-15, 13, 6, -11, -4, 5, -13, 5, 3, 2, 6, -1, 4, 12, -10, -13, -7, -4, -5, 6, 9, -14, 1, -6, 13, 7, -8, 10, -4, 11, -8, -3, 1, 5, -7, 4, -13, -13, -5, -3, 4, -14, 11, -14, 5, -13, -12, 13, -10, -10, -4, -15, 13, 13, -14, 11, -3, -15, 6, 1, 3, 5, 13, -11, -5, -9, 1, -2, -14, 11, 10, 5, 4, -1, 6, -6, -7, 9, -15, -2, 7, -12, -10, 5, -14, 13, -6, -9, 6, 7, 7, -6, -2, -3, -9, 0, -5, 7, 5, -4, -5, -7, -13, 14, 7, 8, -15, 7, -5, -15, -10, 9}))
}
