package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	a, b, c := readInt(), readInt(), readInt()
	if a != b {
		if a > b {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
	} else {
		if c%2 != 0 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
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
