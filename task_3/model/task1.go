package model

type Task1 struct {
}

func (*Task1) Solution(A []int) int {
	min := A[0]
	max := A[0]
	sum := 0
	for _, item := range A {
		sum += item
		if min == item || max == item {
			return 0
		}
		if min > item {
			min = item
		}
		if max < item {
			max = item
		}
	}
	sumRight := (min + max) * (max - min + 1) / 2
	if sumRight != sum {
		return 0
	}
	return 1
}
