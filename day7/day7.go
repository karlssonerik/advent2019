package main

import (
	"fmt"
	"strconv"
)

type result struct {
	res   int
	phase []int
}
type amp struct {
	program    []int
	output     chan int
	input      chan int
	stopped    bool
	lastOutput int
	id         int
}

var result1 result
var result2 result
var nrStopped = 0
var amps []amp

func main() {
	input1 := make(chan int)
	input2 := make(chan int)
	input3 := make(chan int)
	input4 := make(chan int)
	input5 := make(chan int)

	phaseCombos := getNextPhase(5, 9)
	for _, combo := range phaseCombos {
		amps = []amp{
			amp{program: GetInput(),
				output: input2,
				input:  input1,
				id:     1},
			amp{program: GetInput(),
				output: input3,
				input:  input2,
				id:     2},
			amp{program: GetInput(),
				output: input4,
				input:  input3,
				id:     3},
			amp{program: GetInput(),
				output: input5,
				input:  input4,
				id:     4},
			amp{
				program: GetInput(),
				output:  input1,
				input:   input5,
				id:      5},
		}
		runAmps(0, combo)
	}
	fmt.Printf("END!! %d\n", result1)

}

func runAmps(input int, combo []int) {
	nrStopped = 0
	go amps[0].module(combo[0])
	go amps[1].module(combo[1])
	go amps[2].module(combo[2])
	go amps[3].module(combo[3])
	go amps[4].module(combo[4])
	amps[0].input <- input
	stop := false
	for stop == false {
		stop1 := true
		for _, amp := range amps {
			stop1 = stop1 && amp.stopped
		}
		stop = stop1
	}

	if amps[4].lastOutput > result1.res {
		result1.res = amps[4].lastOutput
		result1.phase = combo
	}

	return
}

func (amp *amp) module(phase int) (int, bool) {

	inputIndex := 0
	inp := amp.program

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
		//	fmt.Println("Opcode", opCode)
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
			if inputIndex > 0 {
				inp[inp[i+1]] = <-amp.input
			} else {
				inp[inp[i+1]] = phase
				inputIndex = inputIndex + 1
			}
			i = i + 2
		case "4":
			fmt.Printf("Out: %d\n", inp[inp[i+1]])
			amp.program = inp
			amp.lastOutput = inp[inp[i+1]]
			nextAmp := amp.id % 5
			if nrStopped < 4 && !amps[nextAmp].stopped {
				amp.output <- inp[inp[i+1]]
			}
			i = i + 2

		case "5":
			whatevs := getParam(paramMode1, i+1, inp)
			if whatevs != 0 {
				//	fmt.Println("cond:", whatevs)
				i = getParam(paramMode2, i+2, inp)
				//	fmt.Println("Jumping to ", i)
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
			fmt.Println("AMP stopping", amp.id)
			nrStopped = nrStopped + 1
			amp.program = inp
			amp.stopped = true
			return 0, true
		default:
			fmt.Print("UnkownOPcode\n")
			return -1, true
		}

	}
	fmt.Print("EOF\n")
	return -2, true
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

func getNextPhase(min, max int) [][]int {

	listOfPhases := [][]int{[]int{min}}
	for i := min + 1; i <= max; i++ {
		newSomething := [][]int{}

		for _, combo := range listOfPhases {
			for j := 0; j < (i-min)+1; j++ {
				tail := append([]int{}, combo[j:]...)
				head := append([]int{}, combo[:j]...)
				newSomething = append(newSomething, append(head, append([]int{i}, tail...)...))
			}
		}
		listOfPhases = newSomething
	}

	return listOfPhases
}
