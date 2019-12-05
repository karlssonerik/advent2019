package main

import (
	"fmt"
)

func main() {

	var i, j int
	fmt.Printf("Testing %v\n", GetInput2())
	for i = 0; i < len(GetInput2()); i++ {
		for j = 0; j < len(GetInput2()); j++ {
			res := module(i, j)
			if res == 19690720 {
				fmt.Printf("YEAAH noun: %d verb: %d", i, j)
				return
			}
			// if res > 0 {
			// 	fmt.Printf("meh %d  noun: %d verb: %d\n", res, i, j)
			// }
		}
	}
}

func module(noun, verb int) int {
	//fmt.Printf("Testing noun: %d verb: %d\n", noun, verb)
	inp := GetInput2()

	inp[1] = noun
	inp[2] = verb
	//	fmt.Printf("Testing %v\n", inp)
	var i int
	for i = 0; i < len(inp); i = i + 4 {
		in := inp[i]
		switch in {
		case 1:
			//fmt.Printf("row %d: %d = %d + %d\n", i, inp[i+3], inp[inp[i+1]], inp[inp[i+2]])
			res2 := inp[inp[i+1]] + inp[inp[i+2]]
			inp[inp[i+3]] = res2
		case 2:
			//fmt.Printf("row %d: %d = %d * %d\n", i, inp[i+3], inp[inp[i+1]], inp[inp[i+2]])
			res := inp[inp[i+1]] * inp[inp[i+2]]
			inp[inp[i+3]] = res
		case 99:
			//fmt.Printf("99 at %d! %d", i, inp[0])
			return inp[0]
		default:
			//fmt.Print("LOL!")
			return -1
		}

	}
	return -2
}
