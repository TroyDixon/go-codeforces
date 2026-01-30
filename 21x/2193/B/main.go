package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	a := make([]int, n)
	cnt := 0

	for i := 0; i < n; i++ {
		a[i] = readInt()
	}

	maxval := slices.Max(a)

	for _, val := range a {
		cnt++
		if val == maxval {
			break
		} else {
			continue
		}
	}
	b := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		b[i] = a[i]
	}
	slices.Reverse(b)

	for _, val := range b {
		fmt.Print(val, " ")
	}
	for i := cnt; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

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
