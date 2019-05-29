package main

import (
	"fmt"
	"strings"
)

func myAtoi(str string) int {
	var INT_Max int64 = 1<<31 - 1
	var min_int int64 = 1 << 31
	var INT_MIN int64 = -1 << 31

	if len(str) == 0 || str == "" {
		return 0
	}

	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	var res int64
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
		res = res*10 + int64(str[i]-'0')
		if flag && res >= INT_Max/10 && i < len(str)-1 {
			if (res > INT_Max/10 && str[i+1] >= byte('0') && str[i+1] <= byte('9')) || (res == INT_Max/10 && str[i+1] >= byte('8') && str[i+1] <= byte('9')) {
				return int(INT_Max)
			} else {
				continue
			}
		} else if !flag && res >= min_int/10 && i < len(str)-1 {
			if (res > min_int/10 && str[i+1] >= byte('0') && str[i+1] <= byte('9')) || (res == min_int/10 && str[i+1] >= byte('8') && str[i+1] <= byte('9')) {
				fmt.Println(res)
				return int(INT_MIN)
			} else {
				continue
			}
		}
	}
	fmt.Println("bad")
	if !flag {
		res = res * -1
	}

	return int(res)
}
func main() {
	str := " -1010023630o4"
	fmt.Println(myAtoi(str))
}
