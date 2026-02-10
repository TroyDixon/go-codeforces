package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := readInt()
	b := readStr()
	a := []byte(b)
	zCounter, oCounter := 0, 1

	if n == 1 && a[0] != 49 {
		fmt.Println("1")
		return
	}

	for _, val := range a {
		if val == 48 {
			zCounter++
		} else {
			if zCounter != 0 {
				oCounter += (zCounter / 2) - 1
				zCounter = 0
			} else {
				zCounter = 0
			}
		}
	}
	if a[n-1] == 48 {
		oCounter += (zCounter / 2) - 1
	}

	fmt.Println(oCounter)

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
