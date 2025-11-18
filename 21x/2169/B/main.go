package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {

	// 60 <
	// 62 >
	// 42 *

	s := readStr()
	l := []rune(s)
	n := len(s)
	a := make([]int, n)

	left, right, choose := 0, 0, 0

	for i := 0; i < n; i++ {
		a[i] = int(l[i])
	}

	for _, val := range a {
		if val == 42 {
			choose++
		} else if val == 62 {
			right++
		} else {
			left++
		}
	}

	if choose >= 2 {
		fmt.Println(-1)
		return
	}

	for i := 0; i < n-1; i++ {
		if a[i] != 60 && a[i+1] != 62 {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(choose + max(right, left))

}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		solve()
	}
}

func readInt() int {
	var x int
	fmt.Fscan(in, &x)
	return x
}

func readStr() string {
	var x string
	fmt.Fscan(in, &x)
	return x
}
