package main

import (
	"fmt"
	"strconv"
)

/*
https://www.hackerrank.com/challenges/one-week-preparation-kit-time-conversion/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-one

Given a time in -hour AM/PM format, convert it to military (24-hour) time.

Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.
*/

func timeConversion(s string) string {
	// Write your code here

	timeSuffix := s[len(s)-2:]
	firstTwo := s[0:2]
	firstTwoInt, _ := strconv.Atoi(firstTwo)
	hour := ""

	if timeSuffix == "AM" && firstTwo == "12" {
		hour = "00"
	} else if timeSuffix == "PM" && firstTwoInt < 12 {
		hour = strconv.Itoa(12 + firstTwoInt)
	} else {
		hour = firstTwo
	}

	return hour + s[2:len(s)-2]
}

func main() {
	fmt.Println(timeConversion("08:01:00AM"))
}
