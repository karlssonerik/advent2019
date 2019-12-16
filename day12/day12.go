package main

import (
	"fmt"
	"os"
	"os/exec"
)

type moon struct {
	pos      coordinate
	velocity coordinate
}
type coordinate struct {
	x int
	y int
	z int
}

type moonState struct {
	Io       moon
	Europa   moon
	Ganymede moon
	Callisto moon
}

var moonStates = make(map[moonState]bool)
var repX int64 = 0
var repY int64 = 0
var repZ int64 = 0

/*
<x=-2, y=9, z=-5>
<x=16, y=19, z=9>
<x=0, y=3, z=6>
<x=11, y=0, z=11>
*/
var Io = moon{pos: coordinate{x: -2, y: 9, z: -5}}
var Europa = moon{pos: coordinate{x: 16, y: 19, z: 9}}
var Ganymede = moon{pos: coordinate{x: 0, y: 3, z: 6}}
var Callisto = moon{pos: coordinate{x: 11, y: 0, z: 11}}

// var Io = moon{pos: coordinate{x: -8, y: -10, z: 0}}
// var Europa = moon{pos: coordinate{x: 5, y: 5, z: 10}}
// var Ganymede = moon{pos: coordinate{x: 2, y: -7, z: 3}}
// var Callisto = moon{pos: coordinate{x: 9, y: -8, z: -3}}
var orgs = moonState{
	Io:       moon{pos: coordinate{x: -2, y: 9, z: -5}},
	Europa:   moon{pos: coordinate{x: 16, y: 19, z: 9}},
	Ganymede: moon{pos: coordinate{x: 0, y: 3, z: 6}},
	Callisto: moon{pos: coordinate{x: 11, y: 0, z: 11}},
}

func main() {
	fmt.Println("Step 0")
	fmt.Println("Io", Io.pos, Io.velocity)
	fmt.Println("Europa", Europa.pos, Europa.velocity)
	fmt.Println("Ganymede", Ganymede.pos, Ganymede.velocity)
	fmt.Println("Callisto", Callisto.pos, Callisto.velocity)

	found := false
	i := int64(0)
	for found == false {
		currentIo := Io.pos
		currentEuropa := Europa.pos
		currentGanymede := Ganymede.pos
		currentCallisto := Callisto.pos
		Io.applyGravity(currentEuropa, currentGanymede, currentCallisto)
		Europa.applyGravity(currentIo, currentGanymede, currentCallisto)
		Ganymede.applyGravity(currentEuropa, currentIo, currentCallisto)
		Callisto.applyGravity(currentEuropa, currentGanymede, currentIo)

		// if (i+1)%10 == 0 && i < 1000 {
		// 	fmt.Println("Step ", i+1)
		// 	fmt.Println("Io", Io.pos, Io.velocity)
		// 	fmt.Println("Europa", Europa.pos, Europa.velocity)
		// 	fmt.Println("Ganymede", Ganymede.pos, Ganymede.velocity)
		// 	fmt.Println("Callisto", Callisto.pos, Callisto.velocity)
		// }

		ms := moonState{
			Io:       Io,
			Europa:   Europa,
			Ganymede: Ganymede,
			Callisto: Callisto,
		}

		if checkAxis(ms, 0) && repX == 0 {
			repX = i + 1
			fmt.Println("found x", i)
		}
		if checkAxis(ms, 1) && repY == 0 {
			repY = i + 1
			fmt.Println("found y", i)
		}
		if checkAxis(ms, 2) && repZ == 0 {
			repZ = i + 1
			fmt.Println("found z", i)
		}

		if repX > 0 && repY > 0 && repZ > 0 {
			found = true
			fmt.Println("maybe step", LCM(repX, repY, repZ))
		}

		if i%100000 == 0 {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println("step", i)
		}

		if ms == orgs {
			found = true
			fmt.Println("old state found after step ", i)
		}
		// if (i + 1) == 4686774924 {
		// 	fmt.Println("satan", i, ms, orgs)
		// 	found = true
		// }
		i++
	}

}

func (m *moon) applyVelocity(nextvelocity coordinate) {
	m.pos.x = m.pos.x + m.velocity.x
	m.pos.y = m.pos.y + m.velocity.y
	m.pos.z = m.pos.z + m.velocity.z
	m.velocity = nextvelocity
}

func (m *moon) applyGravity(moon1, moon2, moon3 coordinate) {
	dx := gravityAffectAxis(m.pos.x, moon1.x, moon2.x, moon3.x)
	m.pos.x = m.pos.x + dx
	dy := gravityAffectAxis(m.pos.y, moon1.y, moon2.y, moon3.y)
	m.pos.y = m.pos.y + dy
	dz := gravityAffectAxis(m.pos.z, moon1.z, moon2.z, moon3.z)
	m.pos.z = m.pos.z + dz
	nextvelocity := coordinate{
		x: dx + m.velocity.x,
		y: dy + m.velocity.y,
		z: dz + m.velocity.z,
	}
	m.applyVelocity(nextvelocity)
}

func gravityAffectAxis(orgPos, otherPos1, otherPos2, otherPos3 int) int {
	return (gravityAffectAxisPerMoon(orgPos, otherPos1) + gravityAffectAxisPerMoon(orgPos, otherPos2) + gravityAffectAxisPerMoon(orgPos, otherPos3))
}

func gravityAffectAxisPerMoon(orgPos, otherPos int) int {
	if orgPos == otherPos {
		return 0
	}
	if orgPos > otherPos {
		return -1
	}
	return 1
}

func checkAxis(moons moonState, axis int) bool {
	switch axis {
	case 0:
		return checkOrg(moons.Callisto.pos.x, moons.Io.pos.x, moons.Europa.pos.x, moons.Ganymede.pos.x,
			orgs.Callisto.pos.x, orgs.Io.pos.x, orgs.Europa.pos.x, orgs.Ganymede.pos.x) &&
			checkOrg(moons.Callisto.velocity.x, moons.Io.velocity.x, moons.Europa.velocity.x, moons.Ganymede.velocity.x,
				orgs.Callisto.velocity.x, orgs.Io.velocity.x, orgs.Europa.velocity.x, orgs.Ganymede.velocity.x)
	case 1:
		return checkOrg(moons.Callisto.pos.y, moons.Io.pos.y, moons.Europa.pos.y, moons.Ganymede.pos.y,
			orgs.Callisto.pos.y, orgs.Io.pos.y, orgs.Europa.pos.y, orgs.Ganymede.pos.y) &&
			checkOrg(moons.Callisto.velocity.y, moons.Io.velocity.y, moons.Europa.velocity.y, moons.Ganymede.velocity.y,
				orgs.Callisto.velocity.y, orgs.Io.velocity.y, orgs.Europa.velocity.y, orgs.Ganymede.velocity.y)
	case 2:
		return checkOrg(moons.Callisto.pos.z, moons.Io.pos.z, moons.Europa.pos.z, moons.Ganymede.pos.z,
			orgs.Callisto.pos.z, orgs.Io.pos.z, orgs.Europa.pos.z, orgs.Ganymede.pos.z) &&
			checkOrg(moons.Callisto.velocity.z, moons.Io.velocity.z, moons.Europa.velocity.z, moons.Ganymede.velocity.z,
				orgs.Callisto.velocity.z, orgs.Io.velocity.z, orgs.Europa.velocity.z, orgs.Ganymede.velocity.z)
	}
	return false
}

func checkOrg(v1, v2, v3, v4, o1, o2, o3, o4 int) bool {
	return v1 == o1 &&
		v2 == o2 &&
		v3 == o3 &&
		v4 == o4
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
