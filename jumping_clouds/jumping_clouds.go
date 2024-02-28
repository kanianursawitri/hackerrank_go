package main

import (
	"fmt"
)

// func jumpingOnClouds(c []int32) int32 {
// 	// Write your code here
// 	var steps int32
// 	var prevNumbers string
// 	var justJumped bool

// 	for _, number := range c {
// 		if len(prevNumbers) < 3 {
// 			prevNumbers += strconv.Itoa(int(number))
// 		} else {
// 			prevNumbers = prevNumbers[1:3] + strconv.Itoa(int(number))
// 		}

// 		if number == 0 && (prevNumbers == "010" || prevNumbers == "000") && !justJumped {
// 			justJumped = true
// 		} else {
// 			justJumped = false
// 			steps += 1
// 		}
// 	}

// 	return steps - 1
// }

func jumpingOnClouds(c []int32) int32 {
	n := len(c)
	i := 0
	steps := 0
	for i < n {
		if i+2 < n && c[i+2] == 0 {
			i = i + 2
		} else {
			i = i + 1
		}

		steps += 1
	}
	return int32(steps) - 1
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 1, 0, 0, 0, 0, 0, 0, 0, 0}))
	// fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0}))
}
