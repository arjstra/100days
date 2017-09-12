package main

import "fmt"

type HanoiTower struct {
	disks []int
	rRod  string
	mRod  string
	lRod  string
}

func (t *HanoiTower) assistRod(src string, dst string) string {
	if t.rRod != src && t.rRod != dst {
		return t.rRod
	}
	if t.mRod != src && t.mRod != dst {
		return t.mRod
	}
	return t.lRod
}

func (t *HanoiTower) move(src string, dst string, count int) {
	if count == 1 {
		fmt.Printf("move disk: %d from %s to %s\n", t.disks[0], src, dst)
		return
	}
	mid := t.assistRod(src, dst)
	t.move(src, mid, count-1)
	fmt.Printf("move disk: %d from %s to %s\n", t.disks[count-1], src, dst)
	t.move(mid, dst, count-1)
}

func (t *HanoiTower) Run() {
	if len(t.disks) == 0 {
		return
	}
	t.move(t.rRod, t.lRod, len(t.disks))
}

func New(disks []int, src string, dst string, spare string) *HanoiTower {
	return &HanoiTower{disks, src, spare, dst}
}

func main() {
	disks := []int{1, 2, 4, 5, 7, 9, 10, 13, 17}
	t := New(disks, "right", "left", "middle")
	t.Run()
}
