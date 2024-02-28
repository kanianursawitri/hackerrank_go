package main

import (
	"fmt"
	"sort"
)

/*

https://www.hackerrank.com/test/eoipgdk427n/questions/a8taf02a12a

n is odd

*/

func findMedian(arr []int32) int32 {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr[len(arr)/2]
}

func main() {
	fmt.Println(findMedian([]int32{4, 6, 5, 3, 8, 7, 6}))
}
