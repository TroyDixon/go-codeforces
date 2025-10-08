package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	max, min := -1<<31, 1<<31-1
	b, c := 0, 0
	for i := 0; i < n; i++ {
		a := readInt()
		if a > max {
			max = a
		}
		if a < min {
			min = a
		}
	}
	c = n * max
	b = max + ((n - 1) * min)
	fmt.Println(c - b)
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
