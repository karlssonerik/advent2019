package main

import (
	"fmt"
	"strconv"
)

func main() {

	module(2)
	fmt.Printf("END!! %d\n")

}

func module(input int) (int, bool) {

	//inputIndex := 0
	inpArr := GetInput()
	inp := make(map[int]int)
	for i, v := range inpArr {
		inp[i] = v
	}

	i := 0
	paramOffest := 0
	for i < len(inp) {
		in := inp[i]
		strIn := strconv.Itoa(in)
		opCode := ""
		if len(strIn) > 1 {
			opCode = strIn[len(strIn)-2:]
		} else {
			opCode = "0" + strIn[len(strIn)-1:]
		}

		paramMode1 := "0"
		paramMode2 := "0"
		paramMode3 := "0"
		if len(strIn) >= 3 {
			paramMode1 = strIn[len(strIn)-3 : len(strIn)-2]
			if len(strIn) >= 4 {
				paramMode2 = strIn[len(strIn)-4 : len(strIn)-3]
				if len(strIn) >= 5 {
					paramMode3 = strIn[len(strIn)-5 : len(strIn)-4]
				}
			}
		}
		switch opCode {
		case "01":
			res := getParam(paramMode1, i+1, paramOffest, inp) + getParam(paramMode2, i+2, paramOffest, inp)
			writeParam(paramMode3, i+3, paramOffest, res, inp)
			//inp[inp[i+3]] = res
			i = i + 4
		case "02":
			res := getParam(paramMode1, i+1, paramOffest, inp) * getParam(paramMode2, i+2, paramOffest, inp)
			writeParam(paramMode3, i+3, paramOffest, res, inp)
			i = i + 4
		case "03":
			fmt.Println("LOL", strIn)
			writeParam(paramMode1, i+1, paramOffest, input, inp)
			//inp[inp[i+1]] = input
			i = i + 2
		case "04":
			fmt.Printf("Out: %d\n", getParam(paramMode1, i+1, paramOffest, inp))
			i = i + 2
		case "05":
			whatevs := getParam(paramMode1, i+1, paramOffest, inp)
			if whatevs != 0 {
				i = getParam(paramMode2, i+2, paramOffest, inp)
			} else {
				i = i + 3
			}
		case "06":
			if getParam(paramMode1, i+1, paramOffest, inp) == 0 {
				i = getParam(paramMode2, i+2, paramOffest, inp)
			} else {
				i = i + 3
			}
		case "07":
			if getParam(paramMode1, i+1, paramOffest, inp) < getParam(paramMode2, i+2, paramOffest, inp) {
				writeParam(paramMode3, i+3, paramOffest, 1, inp)
				//inp[inp[i+3]] = 1
			} else {
				writeParam(paramMode3, i+3, paramOffest, 0, inp)
				//inp[inp[i+3]] = 0
			}
			i = i + 4
		case "08":
			if getParam(paramMode1, i+1, paramOffest, inp) == getParam(paramMode2, i+2, paramOffest, inp) {
				writeParam(paramMode3, i+3, paramOffest, 1, inp)
				//	inp[inp[i+3]] = 1
			} else {
				writeParam(paramMode3, i+3, paramOffest, 0, inp)
				//	inp[inp[i+3]] = 0
			}
			i = i + 4
		case "09":
			paramOffest = paramOffest + getParam(paramMode1, i+1, paramOffest, inp)
			i = i + 2
		case "99":
			fmt.Println("AMP stopping")
			return 0, true
		default:
			fmt.Print("UnkownOPcode\n")
			return -1, true
		}

	}
	fmt.Print("EOF\n")
	return -2, true
}

func getParam(mode string, index, offset int, input map[int]int) int {
	switch mode {
	case "0":
		return input[input[index]]
	case "1":
		return input[index]
	case "2":
		return input[offset+input[index]]
	}

	return -1
}

func writeParam(mode string, index, offset, in int, input map[int]int) int {
	switch mode {
	case "0":
		input[input[index]] = in
	case "1":
		input[index] = in
	case "2":
		input[offset+input[index]] = in
	}

	return -1
}
