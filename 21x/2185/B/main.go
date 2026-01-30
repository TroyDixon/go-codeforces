package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	a := make([]int, n)
	maxx := 0
	for i := 0; i < n; i++ {
		a[i] = readInt()
		if a[i] > maxx {
			maxx = a[i]
		}
	}
	ans := maxx * n
	fmt.Println(ans)
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
