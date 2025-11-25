package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	m := make(map[int]int)
	cnt := 0
	size := 0
	for i := 0; i < n; i++ {
		v := readInt()
		m[v]++
	}
	for val := range m {
		if val > size {
			size = val
		}
	}
	for i := 0; i <= size; i++ {
		if m[i] > i {
			cnt += m[i] - i
		} else if m[i] < i {
			cnt += m[i]
		}
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
