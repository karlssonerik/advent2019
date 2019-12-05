package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	res := 0
	for i := 231832; i < 767346; i++ {
		strI := strconv.Itoa(i)
		matched := false
		for j := 0; j < 10; j++ {
			nr := fmt.Sprintf("%d%d", j, j)
			match, _ := regexp.MatchString(nr, strI)

			trip := fmt.Sprintf("%d%d%d", j, j, j)
			matchTrip, _ := regexp.MatchString(trip, strI)

			matched = matched || (match && !matchTrip)
		}
		intSlice := []int{}
		intSlice = IntToSlice(i, intSlice)
		lastDigit := -1
		digitOrder := true
		for _, digit := range intSlice {
			digitOrder = digitOrder && digit >= lastDigit
			lastDigit = digit
		}

		if matched && digitOrder {
			res = res + 1
		}
	}
	fmt.Printf("done %d", res)
}

func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}
