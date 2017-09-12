package main

import "fmt"

type MatrixSet struct {
	matrixs [][2]int
}

func New(m [][2]int) *MatrixSet {
	return &MatrixSet{m}
}

func (m *MatrixSet) MCM(beg int, end int) (int, [2]int) {
	if beg == end {
		return 0, m.matrixs[beg]
	}
	if beg == end-1 {
		times := m.matrixs[beg][0] * m.matrixs[beg][1] * m.matrixs[end][1]
		return times, [2]int{m.matrixs[beg][0], m.matrixs[end][1]}
	}

	t1 := m.matrixs[beg][0] * m.matrixs[beg][1] * m.matrixs[beg+1][1]
	t, result := m.MCM(beg+2, end)
	t1 += t + m.matrixs[beg][0]*result[0]*result[1]

	t2 := 0
	t, result = m.MCM(beg+1, end)
	t2 += t + m.matrixs[beg][0]*result[0]*result[1]

	result = [2]int{m.matrixs[beg][0], m.matrixs[end][1]}
	if t1 > t2 {
		return t2, result
	}
	return t1, result
}

func (m *MatrixSet) Run() int {
	if len(m.matrixs) == 0 {
		return 0
	}
	t, _ := m.MCM(0, len(m.matrixs)-1)
	return t
}

func main() {
	matrix := [][2]int{{10, 20}, {20, 30}, {30, 40}, {40, 5}}
	set := New(matrix)
	fmt.Println("minimum number of multiplications is: ", set.Run())
}
