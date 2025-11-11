package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	a := make([]int, 0, n)
	if n%2 == 0 {
		a = append(a, -1)
		for i := 0; i < (n/2)-1; i++ {
			a = append(a, 3)
			a = append(a, -1)
		}
		a = append(a, 2)
	} else {
		a = append(a, -1)
		for i := 0; i < n/2; i++ {
			a = append(a, 3)
			a = append(a, -1)
		}
	}
	for i := range a {
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
