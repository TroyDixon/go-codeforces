package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, s, flag := readInt(), readStr(), true
	c, same := []rune(s), []int{}

	for i := 0; i < n/2; i++ {
		if c[i] != c[n-i-1] {
			flag = false
			break
		}
	}

	if flag {
		fmt.Println(0)
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			same = append(same, i+1)
		}
	}

	fmt.Println(len(same))

	ans := make([]string, len(same))
	for j, val := range same {
		ans[j] = strconv.Itoa(val)
	}
	fmt.Println(strings.Join(ans, " "))
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
