package main

import (
	"fmt"
)

func convert(s string, numRows int) string {

	if len(s) == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	var internal [][]byte
	var res string
	dif := len(s)%(2*numRows-2) - (numRows - 1)
	if dif <= 0 {
		dif = len(s) % (2*numRows - 2)
	}
	c := len(s)/(2*numRows-2)*(numRows-1) + dif
	internal = make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		internal[i] = make([]byte, c)
	}

	for i := 0; i < len(s); i++ {
		//row
		//colum
		First := i / (numRows + numRows - 2)
		Second := i % (numRows + numRows - 2)
		var row, column int
		if Second-(numRows-1) > 0 {
			row = numRows - 1 - (Second - numRows + 1)
			column = First*(numRows-1) + Second - (numRows - 1)
		} else {
			row = Second
			column = First * (numRows - 1)
		}
		//fmt.Print(internal[row][column])
		internal[row][column] = byte(s[i])
		//fmt.Print(string(s[i]))
	}

	for _, v := range internal {
		for _, v1 := range v {
			if v1 != 0 {
				res = fmt.Sprintf("%s%s", res, string(v1))
			}
		}

	}
	return res
}
func main() {
	fmt.Println("zi")
	s := "LEETCODEISHIRINGajj"
	res := convert(s, 5)
	fmt.Println(res)
}
