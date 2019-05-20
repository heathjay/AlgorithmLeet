package main

import "fmt"

func Print(arr *[3][4]int) {
	//获得行列最大值
	MaxRow := len((*arr)) - 1
	MaxColumn := len((*arr)[0]) - 1
	AR := 0
	AC := 0
	BR := 0
	BC := 0
	flag := false
	for {

		Move(arr, AR, AC, BR, BC, flag)
		if AR == MaxRow && AC == MaxColumn {
			break
		}
		flag = !flag
		if AC == MaxColumn {
			AR++
		} else {
			AC++
		}
		if BR == MaxRow {
			BC++
		} else {
			BR++
		}

	}
}
func Move(arr *[3][4]int, AR, AC, BR, BC int, flag bool) {
	for {
		if AR == BR && AC == BC {
			fmt.Print((*arr)[AR][AC], " ")
			break
		}
		if flag {
			fmt.Print((*arr)[AR][AC], " ")
			AR++
			AC--
		} else {
			fmt.Print((*arr)[BR][BC], " ")
			BC++
			BR--
		}

	}
}
func main() {
	arr := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//1,2,5,9,6,3,4,7,10,11,8,12
	Print(&arr)

}
