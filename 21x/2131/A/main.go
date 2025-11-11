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
	b := make([]int, n)
	var cnt int
	for i := 0; i < n; i++ {
		a[i] = readInt()
	}
	for i := 0; i < n; i++ {
		b[i] = readInt()
	}
	for i := range a {
		cnt += (max(a[i]-b[i], 0))
	}
	fmt.Println(cnt + 1)
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
