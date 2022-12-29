package model

type Task3 struct {
}

func (*Task3) Solution(A []int) int {
	if len(A)%2 != 0 {
		return 0
	}
	result := A[0]
	for i := 1; i < len(A); i++ {
		result ^= A[i] // XOR each element in array
	}
	return result
}
