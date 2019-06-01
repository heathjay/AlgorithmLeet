package main

import (
	"fmt"
	"runtime"
)

func work1(dist *[][]int, x int, y int, tX int, tY int, s1 *string, s2 *string) {

	for i := x; i < x+tY; i++ {
		for j := y; j < y+tX; j++ {
			dif := 1
			if (*s2)[i-1] == (*s1)[j-1] {
				dif = 0
			}
			(*dist)[i][j] = Min(Min((*dist)[i-1][j]+1, (*dist)[i][j-1]+1), (*dist)[i-1][j-1]+dif)
		}
	}

}
func work(dist *[][]int, x int, y int, tX int, tY int, s1 *string, s2 *string, exit chan bool, Final chan bool) {

	for i := x; i < x+tY; i++ {
		for j := y; j < y+tX; j++ {
			dif := 1
			if (*s2)[i-1] == (*s1)[j-1] {
				dif = 0
			}
			(*dist)[i][j] = Min(Min((*dist)[i-1][j]+1, (*dist)[i][j-1]+1), (*dist)[i-1][j-1]+dif)
		}
	}
	exit <- true
	Final <- true
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func MoveGo(cpus int, AR int, AC int, BR int, BC int, dp *[][]int, xB int, yB int, exit chan bool, s1 *string, s2 *string, xFlag bool, yFlag bool, xT int, yT int, Final chan bool) {
	c := 1

	for {
		tX := xT
		tY := yT
		if AR <= 0 || AC <= 0 {
			break
		}

		if !xFlag {
			if AC == cpus*xT+1 {
				tX = xB
			}
		}

		if !yFlag {
			if AR == cpus*yT+1 {
				tY = len(*s2) - cpus*yT
			}
		}

		go work(dp, AR, AC, tX, tY, s1, s2, exit, Final)
		//	fmt.Println("xB=", xB, "yB=", yB, "xT=", xT, "yT=", yT, "tX=", tX, "tY=", tY)
		//	fmt.Println("AR=", AR, "AC=", AC, "BR=", BR, "BC=", BC, "i=", c)
		c++
		if AR == BR && AC == BC {
			break
		}
		AC -= xT
		AR += yT

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

	//fmt.Println(cpus)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = i
	}
	for i := 0; i < m; i++ {
		dp[0][i] = i
	}

	xFlag := false
	yFlag := false
	wC := cpus
	hC := cpus

	if (m-1)%cpus == 0 {
		xFlag = true
	} else {
		wC += 1
	}

	if (n-1)%cpus == 0 {
		yFlag = true
	} else {
		hC += 1
	}
	var count int
	Final := make(chan bool, wC*hC)
	AR := 1
	AC := 1
	BR := 1
	BC := 1

	for {

		xT := (m - 1) / cpus
		yT := (n - 1) / cpus

		xB := xT
		yB := yT

		//channel size
		if yT == 0 {
			work1(&dp, 1, 1, len(s1), len(s2), &s1, &s2)
		} else {
			count = (BR-AR)/yT + 1
		}

		//	fmt.Println("Xt=", xT)
		exitChan := make(chan bool, count)

		//the rest
		if AC == cpus*xT+1 {
			xB = m - 1 - cpus*xT
		}

		if AR == cpus*yT+1 {
			yB = n - 1 - cpus*yT
		}

		MoveGo(cpus, AR, AC, BR, BC, &dp, xB, yB, exitChan, &s1, &s2, xFlag, yFlag, xT, yT, Final)

		//fmt.Println("AR=", AR, "AC=", AC, " BR=", BR, "BC=", BC, " Count=", count)

		// if AC == BC && AR == BR  {
		// 	break
		// }

		if xFlag && yFlag {
			if AC == (cpus-1)*xT+1 && AR == (cpus-1)*yT+1 {
				break
			}
		} else if xFlag && !yFlag {
			if AC == (cpus-1)*xT+1 && AR == cpus*yT+1 {
				break
			}
		} else if !xFlag && yFlag {
			if AC == cpus*xT+1 && AR == (cpus-1)*yT+1 {
				break
			}
		} else if !xFlag && !yFlag {
			if AC == cpus*xT+1 && AR == cpus*yT+1 {
				break
			}
		}

		if xFlag {
			if AC == (cpus-1)*xT+1 {
				AR += yT
			} else {
				AC += xT
			}
		} else {
			if AC == cpus*xT+1 {
				AR += yT
			} else {
				AC += xT
			}
		}

		if yFlag {
			if BR == (cpus-1)*yT+1 {
				BC += xT
			} else {
				BR += yT
			}
		} else {
			if BR == cpus*yT+1 {
				BC += xT
			} else {
				BR += yT
			}
		}
		for i := 0; i < count; i++ {
			_, ok := <-exitChan
			if !ok {
				break
			}
		}
		//exitMChan <- true
		//fmt.Println("AC=", AC, " BC=", BC, " Count=", count)
		//time.Sleep(3 * time.Second)
		//		fmt.Println("close count:", count)
		close(exitChan)
	}
	// res := make([]bool, cap)
	for i := 0; i < hC*wC; i++ {
		<-Final
	}
	close(Final)
	//fmt.Println(len(res))
	//time.Sleep(10 * time.Second)
	// for _, v := range dp {
	// 	fmt.Println(v)

	// }
	return dp[len(s2)][len(s1)]
}
func main() {
	// s1 := "goodjiajfijawijobjobjobjonu"
	// s2 := "jijiaj"
	// fmt.Println(len(s1), " ", len(s2))
	var s1 string
	var s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	res := Distribut(s1, s2)
	fmt.Println(res)
}
