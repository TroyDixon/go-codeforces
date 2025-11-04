// Not the most elegant solution, but it worked

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
	a := make([]int, n)
	b := make([]int, n)
	var arr []int
	flag := false
	for i := 0; i < n; i++ {
		a[i] = readInt()
	}
	copy(b, a)
	sort.Ints(b)
	for k := 0; k < n-1; k++ {
		if a[k]%2 != a[k+1]%2 {
			flag = true
			break
		} else {
			flag = false
		}
	}
	if flag {
		arr = b
	} else {
		arr = a
	}
	for i, val := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
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
