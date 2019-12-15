package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

var score = coordinate{
	x: -1,
	y: 0,
}

var grid = make(map[coordinate]int)

func main() {
	origo := coordinate{x: 0, y: 0}
	grid[origo] = 1
	module()
	printBoard()

	fmt.Printf("END!! %d\n")

}

func module() (int, bool) {

	//inputIndex := 0
	inpArr := getInput()
	inp := make(map[int]int)
	for i, v := range inpArr {
		inp[i] = v
	}

	i := 0
	paramOffest := 0
	outputState := 0
	cacheCoord := coordinate{}
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
			i = i + 4
		case "02":
			res := getParam(paramMode1, i+1, paramOffest, inp) * getParam(paramMode2, i+2, paramOffest, inp)
			writeParam(paramMode3, i+3, paramOffest, res, inp)
			i = i + 4
		case "03":
			printBoard()
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			input, _ := strconv.Atoi(text)
			writeParam(paramMode1, i+1, paramOffest, input, inp)
			i = i + 2
		case "04":
			switch outputState % 3 {
			case 0:
				cacheCoord.x = getParam(paramMode1, i+1, paramOffest, inp)
			case 1:
				cacheCoord.y = getParam(paramMode1, i+1, paramOffest, inp)
			case 2:
				if cacheCoord == score {
					fmt.Println("Score: ", getParam(paramMode1, i+1, paramOffest, inp))
				} else {
					grid[cacheCoord] = getParam(paramMode1, i+1, paramOffest, inp)
				}
			}
			outputState = outputState + 1
			//	fmt.Printf("Out: %d\n", getParam(paramMode1, i+1, paramOffest, inp))
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
			} else {
				writeParam(paramMode3, i+3, paramOffest, 0, inp)
			}
			i = i + 4
		case "08":
			if getParam(paramMode1, i+1, paramOffest, inp) == getParam(paramMode2, i+2, paramOffest, inp) {
				writeParam(paramMode3, i+3, paramOffest, 1, inp)
			} else {
				writeParam(paramMode3, i+3, paramOffest, 0, inp)
			}
			i = i + 4
		case "09":
			paramOffest = paramOffest + getParam(paramMode1, i+1, paramOffest, inp)
			i = i + 2
		case "99":
			fmt.Println("Shuting down stopping")
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

func printBoard() {
	highestX := 0
	highestY := 22
	lowestX := 0
	lowestY := 0
	for k, _ := range grid {
		if k.x > highestX {
			highestX = k.x
		}
		// if k.y > highestY {
		// 	highestY = k.x
		// }
		if k.x < lowestX {
			lowestX = k.y
		}
		if k.y < lowestY {
			lowestY = k.y
		}
	}

	for j := lowestY; j <= highestY; j++ {
		for i := lowestX; i <= highestX; i++ {
			coo := coordinate{
				x: i,
				y: j,
			}
			switch grid[coo] {
			case 0:
				fmt.Print("   ")
			case 1:
				fmt.Print(" | ")
			case 2:
				fmt.Print(" ⌧ ")
			case 3:
				fmt.Print(" - ")
			case 4:
				fmt.Print(" ○ ")
			}
		}
		fmt.Println("\n")
	}
}
