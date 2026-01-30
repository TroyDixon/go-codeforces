package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, s, x := readInt(), readInt(), readInt()
	var sum int
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readInt()
		sum += a[i]
	}

	if sum > s {
		fmt.Println("NO")
	} else if s%x != sum%x {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
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
