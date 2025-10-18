package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	var n, m, i, j int
	n, m, i, j = readInt(), readInt(), readInt(), readInt()
	i++
	j++

	var x1, y1, x2, y2 int

	x1, y1 = 1, 1

	x2, y2 = (n), (m)

	fmt.Println(x1, y1, x2, y2)
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
