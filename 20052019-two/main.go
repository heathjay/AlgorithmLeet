package main

import (
	"fmt"
)

func twoSum(nums *[]int, target int) []int {
	m := make(map[int]int, len(*nums))
	//res := make([]int, 2)
	for i, v := range *nums {
		m[v] = i
	}
	for i, v := range *nums {
		tmp := target - v
		v1, ok := m[tmp]
		if ok && v1 != i {
			// res[0] = i
			// res[1] = v1
			return []int{i, v1}
		}
	}
	//return res
	return []int{-1}
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 26
	res := twoSum(&nums, target)
	fmt.Println(res)
}
