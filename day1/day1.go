package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum int

	for _, in := range Input {
		nr, _ := strconv.Atoi(in)
		for {
			nr = module(nr)
			if nr > 0 {
				sum += nr
			} else {
				break
			}
		}
	}
	fmt.Printf("%d", sum)
}

func module(in int) int {
	return (in / 3) - 2
}
