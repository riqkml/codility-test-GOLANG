package main

func solution(H []int) int {
	var l []int
	head := -1
	stack := 0
	i := 0

	for i < len(H) {
        var h = H[i];
        if head < 0 {
            head++
            l[head] = h
            i++
        } else if l[head] == h {
            i++
        } else if l[head] < h {
            head++
            l[head] = h
            i++
        } else { 
            head--
			stack++
        }
    }
	return stack + head + 1
}

func main() {
	solution([]int{8, 8, 5, 7, 9, 8, 7, 4, 8})
}