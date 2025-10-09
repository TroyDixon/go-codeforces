package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, c, d := readInt(), readInt(), readInt()
	a := make([]int, n*n)
	for i := 0; i < n*n; i++ {
		a[i] = readInt()
	}
	sort.Ints(a)
	b := make([]int, n*n)
	b[0] = a[0]
	for i := 1; i < n; i++ {
		b[i] = b[i-1] + c
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			b[n*i+j] = b[n*(i-1)+j] + d
		}
	}
	sort.Ints(b)
	if slices.Equal(a, b) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
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
