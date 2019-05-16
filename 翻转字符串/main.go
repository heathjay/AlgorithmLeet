package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "lierzi is my son."
	str2 := strings.Split(str, " ")
	fmt.Println(str2)
	for i := 0; i < len(str2); i++ {
		fmt.Print(str2[len(str2)-i-1], " ")
	}
}
