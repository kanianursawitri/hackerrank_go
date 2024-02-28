package main

import "fmt"

/*
https://www.hackerrank.com/challenges/one-week-preparation-kit-mini-max-sum/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-one

Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.
*/

/*

//FIRST ATTEMPT

func sortArr(arr []int32) []int32 {
	for j := 0; j < 5; j++ {
		for i := 0; i < 5-j; i++ {
			if (i+1 < 5) && arr[i] > arr[i+1] {
				temp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp
			}
		}
	}
	return arr
}

func miniMaxSum(arr []int32) {
	// Write your code here

	arr = sortArr(arr)
	var min, max int64
	for i := 0; i < 5; i++ {
		if i < 4 {
			min += int64(arr[i])
		}
		if i > 0 {
			max += int64(arr[i])
		}
	}

	fmt.Println(min, max)
}

*/

func miniMaxSum(arr []int32) {
	// Write your code here
	var min, max, sum int64
	min = int64(arr[0])
	max = int64(arr[0])
	for i := 0; i < 5; i++ {
		curr := int64(arr[i])
		if min >= curr {
			min = curr
		}
		if max <= curr {
			max = curr
		}
		sum += curr
	}

	fmt.Println(sum-max, sum-min)
}

func main() {
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
}
