package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	p1, p2, p3 := readInt(), readInt(), readInt()

	if (p1+p2+p3)%2 != 0 {
		fmt.Println(-1)
		return
	}
	if p3 >= p1+p2 {
		fmt.Println(p1 + p2)
	} else {
		fmt.Println(((p1 + p2 - p3) / 2) + p3)
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
