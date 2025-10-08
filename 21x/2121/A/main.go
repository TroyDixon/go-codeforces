package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	var n, s, cnt int
	n, s = readInt(), readInt()
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = readInt()
	}
	minElement := pos[0]
	maxElement := pos[n-1]
	if s <= minElement {
		cnt = maxElement - s
	} else if s >= maxElement {
		cnt = s - minElement
	} else {
		sTox1 := (s - minElement) + (maxElement - minElement)
		sToxn := (maxElement - s) + (maxElement - minElement)
		cnt = min(sTox1, sToxn)
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
