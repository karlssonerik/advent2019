package main

import (
	"fmt"
	"strings"
)

var directOrbits = make(map[string]string)

//var indirectOrbits = make(map[string]int)

func main() {

	for _, directOrbit := range Input {
		orbs := strings.Split(directOrbit, ")")
		directOrbits[orbs[1]] = orbs[0]
	}
	getOrbitsFromYOUTOSAN()
	getOrbits()

}

func getOrbitsFromYOUTOSAN() {
	res := 0
	sanPath := getAncetsors("SAN")
	youPath := getAncetsors("YOU")
	for i, step := range youPath {
		if i == (len(sanPath)-1) || step != sanPath[i] {
			res = len(youPath[i:]) + len(sanPath[i:])
			break
		}

	}
	fmt.Printf("Orbits to Santa END! Res: %d \n", res)
}

func getAncetsors(orb string) []string {
	ancestors := []string{}
	ancestor := directOrbits[orb]
	for ancestor != "" {
		ancestors = append([]string{ancestor}, ancestors...)
		ancestor = directOrbits[ancestor]
	}
	return ancestors
}

func getOrbits() {
	res := 0
	totIndirectOrbits := 0
	valDiretOrbs := 0
	for _, parent := range directOrbits {
		valDiretOrbs = valDiretOrbs + 1
		ancestor := directOrbits[parent]
		for ancestor != "" {
			totIndirectOrbits = totIndirectOrbits + 1
			ancestor = directOrbits[ancestor]
		}
	}
	res = totIndirectOrbits + valDiretOrbs
	fmt.Printf("Number of OrbitsEND! Res: %d (%d + %d)\n", res, totIndirectOrbits, valDiretOrbs)
}
