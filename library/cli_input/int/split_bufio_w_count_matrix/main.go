package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	matrix := make(map[int]int)
	for i := 0; i < n; i++ {
		x := nextInt()
		matrix[i] = x
	}

	fmt.Println(matrix)
}
