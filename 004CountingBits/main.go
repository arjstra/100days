package main

import "fmt"

func CountBits(n int) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(CountBits(1103)) // 6 bits
}
