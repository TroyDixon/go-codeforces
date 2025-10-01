package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	var n, a, b, c, d, cnt int
	n = readInt()
	for i := 0; i < n; i++ {
		a, b, c, d = readInt(), readInt(), readInt(), readInt()
		if b > d {
			cnt += (a + b - d)
		} else {
			if a > c {
				cnt += (a - c)
			}
		}
	}
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
