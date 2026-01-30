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
	high, low := n, 1
	a := make([]int, n)

	if n%2 == 0 {
		for i := 0; i < n; i += 2 {
			a[i] = high
			a[i+1] = low
			high--
			low++
		}
	}

	if n%2 == 1 {
		a[0] = high
		high--
		for i := 1; i < n; i += 2 {
			a[i] = low
			a[i+1] = high
			high--
			low++
		}
	}
	slices.Reverse(a)

	for _, val := range a {
		fmt.Print(val, " ")
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
