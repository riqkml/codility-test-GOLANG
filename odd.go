package main

func solution() {
	var list = make(map[int]int)

	for _, value := range A {
		list[value]++
	

	for b, value := range list {
		if value % 2 == 1 {
			return b
		}
	}

	return -1
}