package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, s, bestcnt, worstcnt := readInt(), readStr(), 0, 100
	var best, worst rune
	var besti, worsti int
	m := make(map[rune]int)
	a := make([]string, n)
	a = strings.Split(s, "")

	for _, val := range s {
		m[val]++
	}

	for c := range m {
		if m[c] >= bestcnt {
			bestcnt = m[c]
			best = c
		}
	}

	for c := range m {
		if m[c] <= worstcnt {
			worstcnt = m[c]
			worst = c
		}
	}

	for i, val := range s {
		if val == best {
			besti = i
			break
		}
	}

	for i, val := range s {
		if val == worst {
			worsti = i
			break
		}
	}
	if n > 1 {
		if worsti != besti {
			a[worsti] = a[besti]
		} else if worsti+1 < n {
			a[worsti+1] = a[besti]
		} else {
			a[worsti] = a[besti]
		}
	}

	fmt.Println(strings.Join(a, ""))

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
