package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = readInt()
	}

	sorteda := make([]int, n)
	copy(sorteda, a)
	sort.Ints(sorteda)

	if slices.Equal(a, sorteda) {
		fmt.Println("0")
		return
	}

	minVal := a[0]
	maxVal := a[0]

	for _, val := range a {
		if val < minVal {
			minVal = val
		}
	}

	for _, val := range a {
		if val > maxVal {
			maxVal = val
		}
	}
	if minVal == a[n-1] && maxVal == a[0] {
		fmt.Println("3")
	} else if minVal == a[0] || maxVal == a[n-1] {
		fmt.Println("1")
	} else {
		fmt.Println("2")
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
