package main

import (
	"fmt"
)

func main() {
	var bit int
	fmt.Scan(&bit)
	for mask := 0; mask < 1<<uint(bit); mask++ {
		b := make([]bool, bit)
		for i := 0; i < bit; i++ {
			b[i] = 1<<uint(i)&mask > 0
		}
		fmt.Println(b)
	}
}
