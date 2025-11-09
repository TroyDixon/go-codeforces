package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	maxFlag, minFlag := false, false
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readInt()
	}
	x := readInt()
	maxVal, minVal := x, x
	for _, val := range a {
		if val >= maxVal {
			maxFlag = true
		}
		if val <= minVal {
			minFlag = true
		}
	}

	if maxFlag && minFlag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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
