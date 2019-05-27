package main

import (
	"fmt"
	"runtime"
)

func work(dist *[][]int, x int, y int, tX int, tY int, s1 *string, s2 *string, exit chan bool) {
	max := (len((*s1)) + 1) * (len((*s2)) + 1)
	//fmt.Println("tX=", tX, "tY=", tY)
	//fmt.Println(i, " works")
	for {
		//if (*dist)[x-1][y-1] > 0 && (*dist)[x][y-1] > 0 && (*dist)[x-1][y] > 0 {
		//fmt.Println("x=", x, "y=", y)
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

func MoveGo(cpus int, AR int, AC int, BR int, BC int, dp *[][]int, xT int, yT int, exit chan bool, s1 *string, s2 *string) {
	for {
		//go work(&dp, AR, AC
		xB := xT
		yB := yT
		if AR == (cpus-1)*xT+1 {
			xB = len(*s1) - (cpus-1)*xT
		}
		if AC == (cpus-1)*yT+1 {
			yB = len(*s2) - (cpus-1)*yT
		}
		// fmt.Println("AR=", AR, "AC=", AC)
		// fmt.Println("BR=", BR, "BC=", BC)
		go work(dp, AR, AC, xB, yB, s1, s2, exit)
		if AR == BR && AC == BC {
			break
		}
		AR -= xB
		AC += yB
	}
}

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

	fmt.Println("s1=", m)
	fmt.Println("s2=", n)
	fmt.Println("m=", len(dp[0]), "n=", len(dp))
	AR := 1
	AC := 1
	BR := 1
	BC := 1
	count := 0
	for {

		xT := (m - 1) / cpus
		yT := (n - 1) / cpus
		MoveGo(cpus, AR, AC, BR, BC, &dp, xT, yT, exitChan, &s1, &s2)
		if AR == ((cpus-1)*xT+1) && AC == ((cpus-1)*yT+1) {
			break
		}
		if AR == ((cpus-1)*xT + 1) {
			AC += yT
		} else {
			AR += xT
		}

		if BC == ((cpus-1)*yT + 1) {
			BR += xT
		} else {
			BC += yT
		}
		count++
	}

	c := cpus * cpus
	go func() {
		for range exitChan {
			<-exitChan
			c--
			if c == 0 {
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
	s1 := "doggjoas"
	s2 := "goodjob"
	// var s1 string
	// var s2 string
	// fmt.Scanln(&s1)
	// fmt.Scanln(&s2)
	res := Distribut(s1, s2)
	fmt.Println(res)
}
