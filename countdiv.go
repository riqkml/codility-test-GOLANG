package main

import (
	"fmt"
	"math"
)

func solution(A int, B int, K int) int {
	var count int = 0

	if A % K == 0 {
		count = 1
	} 

	mainSide := math.Floor(float64(B / K)) 
	secSide := math.Floor(float64(A / K))
	result := int(mainSide) - int(secSide) + count
	return result
}

func main() { 
	fmt.Println(solution(6,11,3))
}
