package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	a := make([]int, n)
	var count int
	for i := 0; i < n; i++ {
		a[i] = readInt()
	}
	for i := range a {
		if a[i] == 0 {
			count++
		}
	}
	if count < 2 {
		fmt.Println(-1)

	} else {
		fmt.Println(a[0] + a[n-1])

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
