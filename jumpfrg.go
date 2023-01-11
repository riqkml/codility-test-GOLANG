package main

import "fmt"

func Solution(X int, Y int, D int) int {
	dest := Y - X

	if dest % D == 0 {
		return dest / D
	} else {
		return dest / D + 1
	}
}

func main() {
	fmt.Println(Solution(20,100,20))
}