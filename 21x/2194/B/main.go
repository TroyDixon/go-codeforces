package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, x, y := readInt(), readInt(), readInt()
	a := make([]int, n)
	transf, ans := 0, 0

	for i := 0; i < n; i++ {
		a[i] = readInt()
		transf += a[i] / x
	}

	for i := 0; i < n; i++ {
		run := a[i] + y*(transf-(a[i]/x))
		if run > ans {
			ans = run
		}
	}

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
