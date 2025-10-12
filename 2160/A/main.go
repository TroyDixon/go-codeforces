package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	var a int
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = readInt()
	}
	sort.Ints(res)
	for _, v := range res {
		if v == a {
			a++
		} else if v < a {
		} else {
			break
		}
	}
	fmt.Println(a)
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
