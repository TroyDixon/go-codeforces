package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, c0, c1, h, s := readInt(), readInt(), readInt(), readInt(), readStr()
	ones, zeros := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	if c1+h < c0 {
		c0 = c1 + h
	} else if c0+h < c1 {
		c1 = c0 + h
	}
	fmt.Println(ones*c1 + zeros*c0)
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
