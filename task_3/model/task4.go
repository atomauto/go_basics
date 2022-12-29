package model

type Task4 struct {
}

func (*Task4) Solution(A []int) int {
	m := make(map[int]int)
	for _, item := range A {
		m[item] = 1
	}
	var result int
	for i := 1; i <= len(A); i++ {
		if m[i] != 1 {
			result = i
		}
	}
	return result
}
