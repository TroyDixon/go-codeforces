package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	a, b, c := readStr(), readStr(), readStr()
	aCount, bCount, cCount := 0, 0, 0
	for i := 0; i < 3; i++ {
		switch a {
		case ("A"):
			aCount++
		case ("B"):
			bCount++
		case ("C"):
			cCount++
		default:
		}
		switch b {
		case ("A"):
			aCount++
		case ("B"):
			bCount++
		case ("C"):
			cCount++
		default:
		}
		switch c {
		case ("A"):
			aCount++
		case ("B"):
			bCount++
		case ("C"):
			cCount++
		default:
		}
	}
	if aCount == 2 {
		fmt.Println("A")
	}
	if bCount == 2 {
		fmt.Println("B")
	}
	if cCount == 2 {
		fmt.Println("C")
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

func readStr() string {
	var x string
	fmt.Fscan(in, &x)
	return x
}

func readArr(x ...int) (int, []int) {
	var n int
	if len(x) == 0 {
		fmt.Fscan(in, &n)
	} else {
		n = x[0]
	}
	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	return n, v
}

func print(x ...any) {
	fmt.Fprintln(out, x...)
}
