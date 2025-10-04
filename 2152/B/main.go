package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, rk, ck, rd, cd := readInt(), readInt(), readInt(), readInt(), readInt()
	res := 0
	if rk < rd {
		res = max(res, rd)
	} else if rk > rd {
		res = max(res, n-rd)
	} else if ck < cd {
		res = max(res, cd)
	} else if ck > cd {
		res = max(res, n-cd)
	}
	fmt.Println(res)
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
