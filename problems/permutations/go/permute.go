package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var res [][]int
	permutation := make([]int, len(nums))
	visit := make([]bool, len(nums))

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			copiedPermutation := make([]int, len(nums))
			copy(copiedPermutation, permutation)
			res = append(res, copiedPermutation)
		}

		for i, v := range nums {
			if !visit[i] {
				visit[i] = true
				permutation[index] = v
				backtrack(index + 1)
				visit[i] = false
			}
		}
	}

	backtrack(0)
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
