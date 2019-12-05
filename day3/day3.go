package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	wireTrace1 := make(map[int]map[int]bool)
	wireTrace2 := make(map[int]map[int]bool)
	wireStr1 := ""
	wireStr2 := ""
	wire1 := Input1
	wire2 := Input2
	small := 0
	wireStr1 = drawWire(wire1, wireTrace1, wireStr1)
	fmt.Println("New Wire------------")
	wireStr2 = drawWire(wire2, wireTrace2, wireStr2)
	for k, v := range wireTrace1 {
		v2, ok := wireTrace2[k]
		if ok {
			for k2, bbb := range v {
				if bbb == false {
					fmt.Println("whaa")
				}
				v23, _ := v2[k2]
				if v23 {
					if k == 0 && k2 == 0 {
						continue
					}
					fmt.Printf("Intersect at X: %d, Y: %d\n", k, k2)

					coord := fmt.Sprintf("(%d,%d)", k, k2)
					w1I := strings.Index(wireStr1, coord)

					stepsStr1 := wireStr1[:w1I]
					w2I := strings.Index(wireStr2, coord)
					stepsStr2 := wireStr2[:w2I]

					steps2 := strings.Split(stepsStr2, "(")
					steps1 := strings.Split(stepsStr1, "(")

					// if k < 0 {
					// 	k = k * -1
					// }
					// if k2 < 0 {
					// 	k2 = k2 * -1
					// }
					res := len(steps1) + len(steps2)
					if small == 0 {
						small = res
					}
					if res < small {
						small = res
					}
					fmt.Printf("Len %d\n SMAll %d\n", res, small)
				}
			}
		}

	}
	fmt.Printf("Len 222  %d\n", small)
}

func drawWire(wire []string, wireTrace map[int]map[int]bool, wireStr string) string {

	if wireTrace[0] == nil {
		wireTrace[0] = make(map[int]bool)
	}
	col := wireTrace[0]
	col[0] = true
	currentX := 0
	currentY := 0
	for _, step := range wire {
		nr, err := strconv.Atoi(step[1:])
		xWay := false
		yWay := false
		oldX := currentX
		oldY := currentY
		if err != nil {
			fmt.Printf("%v", err)
			fmt.Println("LOL3")
		}
		switch step[0] {
		case 'R':
			xWay = true
			currentX = currentX + nr
		case 'L':
			xWay = true
			currentX = currentX - nr
		case 'D':
			yWay = true
			currentY = currentY - nr
		case 'U':
			yWay = true
			currentY = currentY + nr
		default:
			fmt.Println("LOL")
		}
		if currentX > oldX || currentY > oldY {
			for i := 1; i <= nr; i++ {
				if xWay {
					if wireTrace[oldX+i] == nil {
						wireTrace[oldX+i] = make(map[int]bool)
					}
					wireTrace[oldX+i][currentY] = true
					wireStr += fmt.Sprintf("(%d,%d)", oldX+i, currentY)
				}
				if yWay {
					if wireTrace[currentX] == nil {
						wireTrace[currentX] = make(map[int]bool)
					}
					wireTrace[currentX][oldY+i] = true
					wireStr += fmt.Sprintf("(%d,%d)", currentX, oldY+i)
				}
			}
		} else {
			for i := 1; i <= nr; i++ {
				if xWay {
					if wireTrace[oldX-i] == nil {
						wireTrace[oldX-i] = make(map[int]bool)
					}
					wireTrace[oldX-i][currentY] = true
					wireStr += fmt.Sprintf("(%d,%d)", oldX-i, currentY)
				}
				if yWay {
					wireTrace[currentX][oldY-i] = true
					wireStr += fmt.Sprintf("(%d,%d)", currentX, oldY-i)
				}
			}
		}
	}
	return wireStr
}
