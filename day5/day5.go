package main

import (
	"fmt"
	"strconv"
)

func main() {

	res := module(5)
	fmt.Printf("END!! %d\n", res)

}

func module(input int) int {
	inp := GetInput()

	i := 0
	for i < len(inp) {
		in := inp[i]
		strIn := strconv.Itoa(in)
		opCode := strIn[len(strIn)-1:]
		paramMode1 := "0"
		paramMode2 := "0"
		// paramMode3 := "0"
		if len(strIn) >= 3 {
			paramMode1 = strIn[len(strIn)-3 : len(strIn)-2]
			if len(strIn) >= 4 {
				paramMode2 = strIn[len(strIn)-4 : len(strIn)-3]
				// if len(strIn) >= 5 {
				// 	paramMode3 = strIn[len(strIn)-5 : len(strIn)-4]
				// }
			}
		}

		switch opCode {
		case "1":

			res := getParam(paramMode1, i+1, inp) + getParam(paramMode2, i+2, inp)
			inp[inp[i+3]] = res
			i = i + 4
		case "2":
			res := getParam(paramMode1, i+1, inp) * getParam(paramMode2, i+2, inp)
			inp[inp[i+3]] = res
			i = i + 4
		case "3":
			inp[inp[i+1]] = input
			i = i + 2
		case "4":
			fmt.Printf("Out: %d", inp[inp[i+1]])
			i = i + 2
		case "5":
			if getParam(paramMode1, i+1, inp) != 0 {
				i = getParam(paramMode2, i+2, inp)
			} else {
				i = i + 3
			}
		case "6":
			if getParam(paramMode1, i+1, inp) == 0 {
				i = getParam(paramMode2, i+2, inp)
			} else {
				i = i + 3
			}
		case "7":
			if getParam(paramMode1, i+1, inp) < getParam(paramMode2, i+2, inp) {
				inp[inp[i+3]] = 1
			} else {
				inp[inp[i+3]] = 0
			}
			i = i + 4
		case "8":
			if getParam(paramMode1, i+1, inp) == getParam(paramMode2, i+2, inp) {
				inp[inp[i+3]] = 1
			} else {
				inp[inp[i+3]] = 0
			}
			i = i + 4
		case "9":
			return inp[0]
		default:
			fmt.Print("UnkownOPcode\n")
			return -1
		}

	}
	return -2
}

func getParam(mode string, index int, input []int) int {
	switch mode {
	case "0":
		return input[input[index]]
	case "1":
		return input[index]
	}

	return -1
}
