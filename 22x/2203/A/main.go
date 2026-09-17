package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, m, d := readInt(), readInt(), readInt()

	if m > d {
		fmt.Println(n)
		return
	}

	cnt := (n + (d / m)) / (d/m + 1)

	fmt.Println(cnt)
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
