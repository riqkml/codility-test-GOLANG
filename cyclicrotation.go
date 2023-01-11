package main

import "fmt"

func Solution(A []int, K int) []int {
	min := 0
	max := 100

	if len(A) >= min &&
		len(A) <= max &&
		K >= min && K <= max {

		K := (len(A) - (K % len(A)))

		return append(A[K:], A[0:K]...)
	}
	return A
}

func main() {
	fmt.Println("test solusi : ",Solution([]int{1, 2, 3, 4, 5}, 3))
}