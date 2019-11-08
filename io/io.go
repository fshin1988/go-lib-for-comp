package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = nextInt()
	}
	for i := 0; i < N; i++ {
		if i != 0 {
			fmt.Fprint(wtr, " ")
		}
		fmt.Fprint(wtr, arr[i])
	}
	fmt.Fprintln(wtr)
	_ = wtr.Flush()
}

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
