package main

func SSolution(N int, A []int) []int {
	var newArray = make([]int, N)
	maxNumberCounter := newArray[0]

	for _, v := range A {
		if v >= 1 && v <= N {
			newArray[v-1] += 1
			if newArray[v-1] > maxNumberCounter {
				maxNumberCounter = newArray[v-1]
			}
		}
		if v > N {
			for i := range newArray {
				newArray[i] = maxNumberCounter
			}
		}
	}

	return newArray
}

func main2() {
	start:=time.Start()
	SSolution(5, []int{3, 4, 4, 6, 1, 4, 4})
}