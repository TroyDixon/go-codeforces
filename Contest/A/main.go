package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, k := readInt(), readInt()
	s := readStr()

	prev := -10000000
	res := 0

	for i := 0; i < n; i++ {
		if s[i] == '1' {

			if i-prev >= k {
				res++
			}
			prev = i
		}
	}

	fmt.Println(res)
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

func readStr() string {
	var x string
	fmt.Fscan(in, &x)
	return x
}
