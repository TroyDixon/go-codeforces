package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	a := make([]int, n)
	cnt := 1
	var mxcnt int

	for i := 0; i < n; i++ {
		a[i] = readInt()
	}

	slices.Sort(a)

	if n == 1 || a[0] == a[n-1] {
		fmt.Println("1")
		return
	}

	for i := 0; i < n-1; i++ {
		if a[i+1] == (a[i] + 1) {
			cnt++
		} else if a[i] == a[i+1] {
			continue
		} else {
			mxcnt = cnt
			cnt = 1
		}
		if mxcnt < cnt {
			mxcnt = cnt
		}
	}
	fmt.Println(mxcnt)
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
