package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	max := 0
	left := 0
	for i, v := range s {
		if m != nil {
			v1, ok := m[v]
			if ok {
				if left < v1+1 {
					left = v1 + 1
				}
			}
		}
		m[v] = i
		if max < i-left+1 {
			max = i - left + 1
		}
	}

	return max
}
func main() {
	s := "aab"
	a := lengthOfLongestSubstring(s)
	fmt.Println(a)
}
