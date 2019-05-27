package main

import (
	"fmt"
	"runtime"
)

func work(dist *[][]int, x int, y int, tX int, tY int, s1 *string, s2 *string, exit chan bool) {
	max := -len((*s1)) * len((*s2))
	//fmt.Println(i, " works")
	for {
		//if (*dist)[x-1][y-1] > 0 && (*dist)[x][y-1] > 0 && (*dist)[x-1][y] > 0 {
		if (*dist)[x-1][y-1] != max && (*dist)[x][y-1] != max && (*dist)[x-1][y] != max {
			break
		}
	}

	for i := x; i < x+tX; i++ {
		for j := y; j < y+tY; j++ {
			dif := 1
			if (*s1)[i-1] == (*s2)[j-1] {
				dif = 0
			}
			(*dist)[i][j] = Min(Min((*dist)[i-1][j]+1, (*dist)[i][j-1]+1), (*dist)[i-1][j-1]+dif)
		}
	}
	exit <- true
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// func MoveGo(cpus int, AR int, AC int, BR int, BC int, dp *[][]int, xT int, yT int, exit chan bool, s1 string, s2 string) {
// 	for {
// 		go work(&dp,AR,AC)
// 		if AR == BR && AC == BC {

// 		}
// 	}
// }

func Distribut(s1, s2 string) int {

	m := len(s1) + 1
	n := len(s2) + 1
	cpus := runtime.NumCPU()

	if n == 1 {
		return len(s1)
	}
	if m == 1 {
		return len(s2)
	}

	exitChan := make(chan bool, cpus*cpus)
	fmt.Println(cpus)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = i
	}
	for i := 0; i < m; i++ {
		dp[0][i] = i
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			dp[i][j] = len((s1)) * len((s2))
		}
	}

	// AR := 1
	// AC := 1
	// BR := 1
	// BC := 1
	// for {
	// 	x := (m - 1) / cpus
	// 	y := (n - 1) / cpus
	// 	MoveGo(cpus, AR, AC, BR, AC, &dp, x, y, exitChan)
	// }

	for i := 0; i < cpus; i++ {
		//first element
		for j := 0; j < cpus; j++ {
			x := (m - 1) / cpus
			y := (n - 1) / cpus

			firstX := x*i + 1
			firstY := y*j + 1

			if i == cpus-1 {
				x = m - firstX
			}

			if j == cpus-1 {
				y = n - firstY
			}
			go work(&dp, firstX, firstY, x, y, &s1, &s2, exitChan)
		}
	}
	count := cpus * cpus
	go func() {
		for range exitChan {
			<-exitChan
			count--
			if count == 0 {
				close(exitChan)
			}
		}

	}()
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[len(s2)][len(s1)]
}
func main() {
	s1 := "doggjoba"
	s2 := "goodjobs"
	// var s1 string
	// var s2 string
	// fmt.Scanln(&s1)
	// fmt.Scanln(&s2)
	res := Distribut(s1, s2)
	fmt.Println(res)
}
