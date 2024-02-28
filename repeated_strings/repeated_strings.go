package main

import "fmt"

/*

link: https://www.hackerrank.com/challenges/repeated-string/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup

There is a string, s, of lowercase English letters that is repeated infinitely many times. Given an integer, n, find and print the number of letter a's in the first  letters of the infinite string.

Constraints

1 <= |s| <= 100
1 <= n <= 10^12

*/

func repeatedString(s string, n int64) int64 {
	// Write your code here
	multiplier := int(n) / len(s)
	remaining := int(n) % len(s)
	numA := 0
	numAremaining := 0

	for idx, letter := range s {
		if string(letter) == "a" {
			numA += 1
			if idx < remaining {
				numAremaining += 1
			}
		}
	}

	result := multiplier*numA + numAremaining

	return int64(result)

}

func main() {
	fmt.Println(repeatedString("aba", 10))
}
