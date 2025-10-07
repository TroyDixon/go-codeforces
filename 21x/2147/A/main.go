package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	var x, y, cnt int
	cnt = 0
	x, y = readInt(), readInt()
	if x < y {
		cnt = 2
	} else if x > y+1 {
		cnt = 3
	} else {
		cnt = -1
	}
	fmt.Println(cnt)
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
