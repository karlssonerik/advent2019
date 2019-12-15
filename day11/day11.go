package main

import (
	"fmt"
	"strconv"
)

type robot struct {
	facing int
	x      int
	y      int
}

type coordinate struct {
	x int
	y int
}

var c3p0 = robot{}
var grid = make(map[coordinate]int)
var grid2 = make(map[coordinate]bool)

func main() {
	origo := coordinate{x: 0, y: 0}
	grid[origo] = 1
	module()
	highestX := 0
	highestY := 0
	lowestX := 0
	lowestY := 0
	for k, _ := range grid {
		if k.x > highestX {
			highestX = k.x
		}
		if k.y > highestY {
			highestY = k.x
		}
		if k.x < lowestX {
			lowestX = k.y
		}
		if k.y < lowestY {
			lowestY = k.y
		}
	}
	painted := 0
	for j := lowestY; j <= highestY; j++ {
		for i := lowestX; i <= highestX; i++ {
			coo := coordinate{
				x: i,
				y: j,
			}
			if grid[coo] == 0 {
				fmt.Print("   ")
			} else {
				fmt.Print(" ⌧ ")
			}
			if grid2[coo] {
				painted = painted + 1
			}
		}
		fmt.Println("\n")
	}

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
	outputTurn := false
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
			coo := coordinate{
				x: c3p0.x,
				y: c3p0.y,
			}
			input := grid[coo]
			writeParam(paramMode1, i+1, paramOffest, input, inp)
			i = i + 2
		case "04":
			if !outputTurn {
				c3p0.paint(getParam(paramMode1, i+1, paramOffest, inp))
				outputTurn = true
			} else {
				c3p0.turn(getParam(paramMode1, i+1, paramOffest, inp))
				c3p0.move()
				outputTurn = false
			}
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

func (r *robot) move() {
	switch r.facing {
	case 0:
		r.y = r.y + 1
	case 1:
		r.x = r.x + 1
	case 2:
		r.y = r.y - 1
	case 3:
		r.x = r.x - 1
	}

	//fmt.Println("on ", r.x, r.y)
}

func (r *robot) turn(way int) {
	switch way {
	case 0:
		if r.facing == 0 {
			r.facing = 3
		} else {
			r.facing = (r.facing - 1)
		}
	case 1:
		r.facing = (r.facing + 1) % 4
	}
}

var paintNo = 0

func (r *robot) paint(color int) {
	coo := coordinate{
		x: c3p0.x,
		y: c3p0.y,
	}

	grid[coo] = color
	// if !grid2[coo] {
	// 	paintNo = paintNo + 1
	// 	fmt.Println("painting no", paintNo)
	// }
	grid2[coo] = true
}
