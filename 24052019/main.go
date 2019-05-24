package main

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	//mar
	if s == "" || len(s) == 0 {
		return ""
	}
	m := make([]byte, 0)
	red := byte('#')
	b := []byte(s)
	// a := b[0:2]
	// fmt.Println(a)
	for _, v := range b {
		m = append(m, red)
		m = append(m, v)

	}
	m = append(m, red)
	fmt.Println(string(m))

	max := 0
	start := 0
	end := 0
	for i := 0; i < len(m); i++ {
		len1 := 0
		for j := 0; j <= i && j < (len(m)-i); j++ {
			if m[i-j] != m[i+j] {
				break
			} else if m[i-j] == m[i+j] {
				len1++
			}
		}
		//fmt.Println(max)
		max = int(math.Max(float64(len1-1), float64(max)))
		if max > (end - start + 1) {
			end = (i-1)/2 + max/2
			start = (i-1)/2 - (max-1)/2

		}
	}
	fmt.Println(max)
	fmt.Println(start)
	fmt.Println(end)
	return string(b[start : end+1])

}
func main() {
	fmt.Println("回文")
	//s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s := "bb"
	res := longestPalindrome(s)
	fmt.Println(res)
}
