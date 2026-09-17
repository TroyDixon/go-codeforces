package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	s := readStr()
	first := 0

	if n%2 == 1 {
		if s[0] == 'b' {
			fmt.Println("NO")
			return
		}
		first++
	}

	for i := first; i < n-1; i += 2 {

		if s[i] != '?' && s[i+1] != '?' && s[i] == s[i+1] {
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

func readStr() string {
	var x string
	fmt.Fscan(in, &x)
	return x
}
