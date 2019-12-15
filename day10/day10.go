package main

import (
	"fmt"
	"math"
	"sort"
)

type coordinate struct {
	x int
	y int
}

// type mapOfUniverse struct {
// 	height    int
// 	width     int
// 	asteroids map[coordinate]bool
// }

// var currentMap = mapOfUniverse{
// 	asteroids: make(map[coordinate]bool),
// }
var asteroids = make(map[coordinate]bool)
var line = coordinate{
	x: -1,
	y: 1,
}

func main() {

	for y, row := range input {
		for x, char := range row {
			if char == '#' {
				ass := coordinate{
					x: x,
					y: y,
				}
				asteroids[ass] = true
			}
		}
	}
	maxC := 0
	ass := coordinate{}
	for k, _ := range asteroids {
		// k := coordinate{
		// 	x: 5,
		// 	y: 8,
		// }

		res := getVisibleAstroids(k)
		if res > maxC {
			maxC = res
			ass = k
		}
		//	fmt.Printf("Ass:%d,%d sees:%d \n", k.x, k.y, res)
	}
	if ass.x == 0 {

	}
	fmt.Printf("Ass:%d,%d sees:%d \n", ass.x, ass.y, maxC)

}

func getVisibleAstroids(from coordinate) int {
	visibleAss := 0
	assOnLine := make(map[float64][]coordinate)
	for k, v := range asteroids {

		if k == from || !v {
			continue
		}

		relativeX := k.x - from.x
		relativeY := k.y - from.y

		assAngle := math.Atan2(float64(relativeY), float64(relativeX)) - math.Atan2(float64(line.y), float64(line.x))
		_, ok := assOnLine[assAngle]
		if !ok {
			visibleAss = visibleAss + 1
		}
		assOnLine[assAngle] = append(assOnLine[assAngle], k)
	}

	if from.x == 37 && from.y == 25 {
		vs := []float64{}
		for k, v := range assOnLine {
			vs = append(vs, k)
			sort.Sort(coordinates(v))
		}
		sort.Float64s(vs)

		vs2 := []float64{}
		closetStart := float64(1000)
		closeStartI := 0
		for i, vss := range vs {
			asses := assOnLine[vss]
			if distance(asses[0]) < float64(closetStart) && asses[0].x == 37 {
				closetStart = distance(asses[0])
				closeStartI = i
			}
		}
		vs2 = append(vs2, vs[closeStartI:]...)
		vs2 = append(vs2, vs[:closeStartI]...)
		vs = vs2
		for i, vss := range vs {
			asses := assOnLine[vss]
			fmt.Println("SKIT  ", i, asses, vss)
		}

		flatlist := []coordinate{}
		asd := false
		index1 := 0
		for asd == false {
			for _, vss := range vs {
				if index1 >= len(assOnLine[vss]) {
					continue
				}
				asses := assOnLine[vss]
				fmt.Printf("twerq %d %v\n", len(flatlist)+1, asses[index1])
				flatlist = append(flatlist, asses[index1])
			}
			index1 = index1 + 1
			asd = index1 == 345
		}
		fmt.Printf("asddLOL %v\n", flatlist[199])

	}

	return visibleAss
}

type coordinates []coordinate

func (a coordinates) Len() int           { return len(a) }
func (a coordinates) Less(i, j int) bool { return distance(a[i]) < distance(a[j]) }
func (a coordinates) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func distance(a coordinate) float64 {
	x := a.x - 11
	y := a.y - 13
	return math.Sqrt(float64(x*x) + float64(y*y))
}
