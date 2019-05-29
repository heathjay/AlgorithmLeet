package main

import (
	"fmt"
	"strings"
)

func myAtoi(str string) int {
	INT_Max := 1<<31 - 1
	INT_MIN := -1 << 31

	if len(str) == 0 || str == "" {
		return 0
	}

	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	var res int
	flag := true

	for i := 0; i < len(str); i++ {
		if i == 0 && str[i] == '-' {
			flag = false
			continue
		} else if i == 0 && str[i] == '+' {
			continue
		}
		if str[i] < byte('0') || str[i] > byte('9') {
			break
		}

		res = res*10 + int(str[i]-'0')
	}

	if flag == false {
		res = -1 * res
	}
	if res < INT_MIN {
		res = INT_MIN
	}
	if res > INT_Max {
		res = INT_Max
	}

	return res
}
func main() {
	str := "+1"
	fmt.Println(myAtoi(str))
}
