package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readInt()
	}

	if a[0] != -1 && a[n-1] != -1 {
		fmt.Println(abs(a[n-1] - a[0]))
		for i := 0; i < n; i++ {
			if a[i] == -1 {
				a[i] = 0
			}
		}
		for i, val := range a {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	} else if a[0] == -1 && a[n-1] != -1 {
		a[0] = a[n-1]
		for i := 0; i < n; i++ {
			if a[i] == -1 {
				a[i] = 0
			}
		}
		fmt.Println(0)
		for i, val := range a {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	} else if a[0] != -1 && a[n-1] == -1 {
		a[n-1] = a[0]
		for i := 0; i < n; i++ {
			if a[i] == -1 {
				a[i] = 0
			}
		}
		fmt.Println(0)
		for i, val := range a {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	} else {
		a[0] = 0
		a[n-1] = 0
		for i := 1; i < n-1; i++ {
			if a[i] == -1 {
				a[i] = 0
			}
		}
		fmt.Println(0)
		for i, val := range a {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
