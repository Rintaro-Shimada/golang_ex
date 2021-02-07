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

func nextString() string {
	sc.Scan()
	s := sc.Text()
	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	matrix := make(map[int]string)
	for i := 0; i < n; i++ {
		x := nextString()
		matrix[i] = x
	}

	fmt.Println(matrix)
}
