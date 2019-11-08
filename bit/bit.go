package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for bit := 0; bit < 1<<uint(N); bit++ {
		arr := make([]bool, N)
		for i := 0; i < N; i++ {
			arr[i] = 1<<uint(i)&bit > 0
		}
		fmt.Println(arr)
	}
}
