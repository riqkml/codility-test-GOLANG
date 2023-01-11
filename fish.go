package main

func solution(A []int, B []int) int {
	generateAliveFish := []int{A[0]}
	sizeUpFish := 0
	currentDirection := B[0]

	for i := 0; i <= len(A)-1; i++ {
		if B[i] == 1 {
			sizeUpFish = A[i]
			generateAliveFish = append(generateAliveFish, A[i])

			currentDirection = 1
		} else if sizeUpFish != 0 && B[i] == 0 {
			if sizeUpFish < A[i] {
				if currentDirection == 1 {
					generateAliveFish[len(generateAliveFish)-1] = A[i]
					currentDirection = 0
				} else if currentDirection == 0 {
					generateAliveFish = append(generateAliveFish, A[i])
				}
			}
		}
	}
	// fmt.Println("index", len(generateAliveFish), generateAliveFish)
	return len(generateAliveFish)
}

func main() {
	solution([]int{4, 3}, []int{1, 1})
}