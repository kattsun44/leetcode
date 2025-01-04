package main

import "fmt"

func maxArea(height []int) int {
	var res int
	l, r := 0, len(height)-1
	for l != r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}
	}
	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
