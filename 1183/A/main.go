package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	var n int
	for n = readInt(); ; n++ {
		temp := 0
		for j := n; j > 0; j /= 10 {
			temp += j % 10
		}
		if temp%4 == 0 {
			fmt.Println(n)
			return
		}
	}
}

func main() {
	defer out.Flush()
	//var t int
	//for fmt.Fscanln(in, &t); t > 0; t-- {
	solve()
}

func readInt() int {
	var x int
	fmt.Fscan(in, &x)
	return x
}
