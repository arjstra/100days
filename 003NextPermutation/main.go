package main

import "fmt"

type Permutation struct {
	items []int
	size  int
}

func (p *Permutation) Permute(ordered []int, count int, left chan int) {
	if count == p.size {
		fmt.Println(ordered)
		return
	}
	for i := 0; i < p.size-count; i++ {
		n := <-left
		if i != 0 && n == ordered[count] {
			left <- n
			continue
		}
		ordered[count] = n
		p.Permute(ordered, count+1, left)
		left <- n
	}
}

func (p *Permutation) Run() {
	if p.size == 0 {
		return
	}

	ch := make(chan int, p.size)
	for _, i := range p.items {
		ch <- i
	}

	p.Permute(make([]int, p.size), 0, ch)

}

func New(items []int) *Permutation {
	return &Permutation{items, len(items)}
}

func main() {
	p := New([]int{1, 2, 3, 3})
	p.Run()
}
