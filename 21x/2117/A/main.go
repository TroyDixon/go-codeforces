package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	var n, x, d, first, last int
	n, x = readInt(), readInt()
	first, last = n+1, -1
	for i := 0; i < n; i++ {
		d = readInt()
		if d == 1 {
			first = min(first, i)
			last = max(last, i)
		}
	}
	if x >= last-first+1 {
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
