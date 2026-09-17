package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, c, k := readInt(), readInt(), readInt()
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = readInt()
	}
	sort.Ints(a)

	for _, val := range a {
		if val > c {
			continue
		}
		diff := c - val
		if diff <= k {
			c += val + diff
			k -= diff
		} else {
			c += val + k
			k = 0
		}
	}
	fmt.Println(c)
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
