package model

type Task2 struct {
}

func (*Task2) Solution(A []int, K int) []int {
	var rotated []int
	if K > len(A) {
		K = K % len(A)
	}
	shift := len(A) - K
	rotated = append(rotated, A[shift:]...)
	rotated = append(rotated, A[:shift]...)
	return rotated
}
