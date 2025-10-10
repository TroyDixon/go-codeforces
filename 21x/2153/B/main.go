package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	x, y, z := readInt(), readInt(), readInt()
	var a, b, c int
	a = x | z
	b = x | y
	c = y | z

	if a&b == x && b&c == y && a&c == z {
		fmt.Println("YES")
	} else if a&b != x {
		fmt.Println("NO")
	} else if b&c != y {
		fmt.Println("NO")
	} else if a&c != z {
		fmt.Println("NO")
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
