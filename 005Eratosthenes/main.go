package main

import "fmt"

func Eratosthenes(n int) {
	if n <= 1 {
		return
	}
	limit := (n+1)>>1
	array := make([]int, n+1)
	array[0], array[1] = 1, 1

	for i:=2 ; i < limit; i++ {
		for j:=2;; j++ {
           idx := i*j
           if idx <= n {
           	   array[idx] = 1
           }else {
	           break
           }
		}
	}
	for i, v := range array {
		if v != 1 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}

func main() {
    Eratosthenes(100000)
}