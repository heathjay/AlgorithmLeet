package main

import "fmt"

func reverse(x int) int {
	intMax := 1<<31 - 1
	intMin := -1 << 31
	rev := 0
	for {
		if x == 0 {
			break
		}
		pop := x % 10
		x /= 10
		if rev > intMax/10 || (rev == intMax/10 && pop > 7) {
			return 0
		}
		if rev < intMin/10 || (rev == intMin/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

func main() {
	res := reverse(390)
	fmt.Println(res)
}
