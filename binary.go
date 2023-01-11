package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getSliceIndex(base []string, predict string)int{
	for i, v := range base {
		if v == predict {
			return i
		}
	}
	return -1
}

func calcGaps(Binary []string, gap int) int{
	firstOne := getSliceIndex(Binary,"1")
	
	if firstOne > -1 {
		firstSearchBinaryArray := Binary[firstOne+1:]

		if firstOne > 0 && firstOne > gap {
			gap = firstOne
		}
		
		if len(firstSearchBinaryArray) > 0 {
			return calcGaps(firstSearchBinaryArray,gap)
		}
	} 
	if gap > 0 {
		return gap
	} else {
		return 0 
	}
	
}

func Solution(N int64) int {

	if N >= 1 && N <= 2147483647 {
		Binary := strings.Split(strconv.FormatInt(int64(N), 2), "")
		return calcGaps(Binary,0)
	}
	
	return 0

	
}

func main() {
	fmt.Println(Solution(9))
	fmt.Println(Solution(1041))
	fmt.Println(Solution(529))
	fmt.Println(Solution(51272))
	fmt.Println(Solution(15))
	fmt.Println(Solution(53))
	fmt.Println(Solution(2147483647))
	fmt.Println(Solution(2049))
}
