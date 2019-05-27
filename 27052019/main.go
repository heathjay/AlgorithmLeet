package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	var internal [][]byte
	internal = make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		internal[i] = make([]byte, 0)
	}

	for i := 0; i < len(s); i++ {
		//row
		First := i / (numRows + numRows - 2)
		Second := i % (numRows + numRows - 2)

		if Second-(numRows-1) > 0 {

		}
	}
}
func main() {
	fmt.Println("zi")
}
