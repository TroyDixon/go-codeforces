package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	x, y := readInt(), readInt()
	flag := true
	delta := (x - 2*y)
	if delta%3 != 0 {
		fmt.Println("NO")
		return
	}
	if y < 0 {
		flag = x >= -4*y
	} else {
		flag = x >= 2*y
	}

	if flag {
		fmt.Println("YES")
	} else {
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
