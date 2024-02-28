package main

import (
	"fmt"
)

/*

https://www.hackerrank.com/challenges/one-week-preparation-kit-diagonal-difference/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-two

Given a square matrix, calculate the absolute difference between the sums of its diagonals.

For example, the square matrix  is shown below:
1 2 3
4 5 6
9 8 9

The left-to-right diagonal = 1+5+9 = 15. The right to left diagonal = 3 + 5 + 9 = 17. Their absolute difference is |15-17|=2

*/

/*

//FIRST ATTEMPT

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	size := len(arr)
	var sumLeftToRight, sumRightToLeft int32
	for line, arrayOut := range arr {
		for idx, elemIn := range arrayOut {
			if idx == line {
				sumLeftToRight += elemIn
			}
			if idx == size-line-1 {
				sumRightToLeft += elemIn
			}
		}
	}

	result := sumLeftToRight - sumRightToLeft
	if result < 0 {
		return -1 * result
	}

	return result
}

1 2 3 4
5 6 7 8
1 2 3 4
5 6 7 8

*/

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	size := len(arr)
	isEven := size%2 == 0
	midRow := size / 2

	var sumLeftToRight, sumRightToLeft int32

	for line, arrayOut := range arr {
		if line != midRow || isEven {
			sumLeftToRight += arrayOut[line]
			sumRightToLeft += arrayOut[size-line-1]
		} else {
			sumLeftToRight += arrayOut[line]
			sumRightToLeft += arrayOut[line]
		}

	}

	result := sumLeftToRight - sumRightToLeft
	if result < 0 {
		return -1 * result
	}

	return result
}

func main() {
	fmt.Println(diagonalDifference([][]int32{{6, 8}, {-6, 9}}))
}
