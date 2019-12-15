package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var layers = make(map[int]string)
var layersAsInts = make(map[int][]int)

const width = 25
const height = 6

const area = width * height

var numberOfLayers = len(Input) / area

func main() {

	for i := 0; i < numberOfLayers; i++ {
		currentLayer := ""
		for j := i * area; j < (i+1)*area; j++ {
			currentLayer += strconv.Itoa(Input[j])
		}
		layers[i] = currentLayer

		currentLayersAsInt := []int{}
		currentLayersAsInt = append(currentLayersAsInt, Input[i*area:(i+1)*area]...)
		layersAsInts[i] = currentLayersAsInt
	}
	printImage()
	fmt.Println("END! ")
}

func printImage() {
	var imageSlice [area]string

	for i := 0; i < numberOfLayers; i++ {
		currentLayer := layersAsInts[i]
		for j := 0; j < area; j++ {
			pixel := strconv.Itoa(currentLayer[j])
			if pixel == "0" {
				pixel = " "
			}
			if imageSlice[j] == "" || imageSlice[j] == "2" {
				imageSlice[j] = pixel
			}
		}
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		for k := 0; k < height; k++ {
			fmt.Println("", imageSlice[k*width:(k+1)*width])
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func checkValidImage() {
	minCount := area
	zeroLayer := ""
	for _, v := range layers {
		if strings.Count(v, "0") < minCount {
			minCount = strings.Count(v, "0")
			zeroLayer = v
		}
	}
	ones := strings.Count(zeroLayer, "1")
	twos := strings.Count(zeroLayer, "2")
	fmt.Println("END!", (ones * twos))
}
