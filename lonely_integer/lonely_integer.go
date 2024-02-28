package main

import (
	"fmt"
)

func lonelyinteger(a []int32) int32 {
	// Write your code here
	mapFound := make(map[int32]int, 0)
	for _, num := range a {
		mapFound[num] += 1
	}

	var result int32
	for key, val := range mapFound {
		if val == 1 {
			result = key
			break
		}
	}

	return result
}

func main() {
	fmt.Println(lonelyinteger([]int32{1, 2, 3, 4, 3, 2, 1}))
}
