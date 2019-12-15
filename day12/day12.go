package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
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

var c = cache.New(30*time.Minute, 40*time.Minute)

var moonStates = make(map[moonState]bool)

/*
<x=-2, y=9, z=-5>
<x=16, y=19, z=9>
<x=0, y=3, z=6>
<x=11, y=0, z=11>
*/
// var Io = moon{pos: coordinate{x: -2, y: 9, z: -5}}
// var Europa = moon{pos: coordinate{x: 16, y: 19, z: 9}}
// var Ganymede = moon{pos: coordinate{x: 0, y: 3, z: 6}}
// var Callisto = moon{pos: coordinate{x: 11, y: 0, z: 11}}

var Io = moon{pos: coordinate{x: -8, y: -10, z: 0}}
var Europa = moon{pos: coordinate{x: 5, y: 5, z: 10}}
var Ganymede = moon{pos: coordinate{x: 2, y: -7, z: 3}}
var Callisto = moon{pos: coordinate{x: 9, y: -8, z: -3}}

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

		// Get the string associated with the key "foo" from the cache
		_, found := c.Get(fmt.Sprintf("%v", ms))

		// found = moonStates[ms]

		if found {
			found = true
			fmt.Println("old state found after step ", i)
		}
		c.Set(fmt.Sprintf("%v", ms), true, cache.NoExpiration)
		// moonStates[ms] = true

		// if i%100000 == 0 {
		// 	cmd := exec.Command("clear") //Linux example, its tested
		// 	cmd.Stdout = os.Stdout
		// 	cmd.Run()
		// 	fmt.Println("step", i)
		// }
		if i > 4686774924 {
			fmt.Println("satan", i)
			found = true
		}
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
