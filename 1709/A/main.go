package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	res := make(map[string]int)
	res["1"] = readInt()
	res["2"] = readInt()
	res["3"] = readInt()
	for i := 0; i < 2; i++ {
		nstr := strconv.Itoa(n)
		n = res[nstr]
		if n == 0 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
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
