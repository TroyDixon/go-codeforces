package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	x, n, m := (readInt()), readInt(), readInt()
	for x > (x/2)+10 && n > 0 {
		n--
		x = (x / 2) + 10
	}
	if x > (m * 10) {
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
