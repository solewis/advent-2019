package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := moon{point{7, 10, 17}, point{0, 0, 0}}
	m2 := moon{point{-2, 7, 0}, point{0, 0, 0}}
	m3 := moon{point{12, 5, 12}, point{0, 0, 0}}
	m4 := moon{point{5, -8, 6}, point{0, 0, 0}}
	moons := []*moon{&m1, &m2, &m3, &m4}
	fmt.Printf("Part 1: %d\n", calcTotalEnergyAfterSteps(moons, 1000))

	moonsP2 := []*moon{&m1, &m2, &m3, &m4}
	fmt.Printf("Part 2: %d\n", calcStepsUntilRepeat(moonsP2))
}

type point struct{ x, y, z int }
type moon struct{ position, velocity point }

func (m *moon) applyGravity(o moon) {
	if m.position.x < o.position.x {
		m.velocity.x++
	} else if m.position.x > o.position.x {
		m.velocity.x--
	}
	if m.position.y < o.position.y {
		m.velocity.y++
	} else if m.position.y > o.position.y {
		m.velocity.y--
	}
	if m.position.z < o.position.z {
		m.velocity.z++
	} else if m.position.z > o.position.z {
		m.velocity.z--
	}
}

func (m *moon) applyGravitySinglePlane(o *moon, getPos func(m *moon) int, setVel func(m *moon, v int)) {
	if getPos(m) < getPos(o) {
		setVel(m, 1)
	} else if getPos(m) > getPos(o) {
		setVel(m, -1)
	}
}

func (m *moon) move() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}

func (m *moon) totalEnergy() int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	pot := abs(m.position.x) + abs(m.position.y) + abs(m.position.z)
	kin := abs(m.velocity.x) + abs(m.velocity.y) + abs(m.velocity.z)
	return pot * kin
}

func calcTotalEnergyAfterSteps(moons []*moon, steps int) int {
	moonPairs := getMoonPairs(moons)
	for i := 0; i < steps; i++ {
		for _, pair := range moonPairs {
			m1 := pair[0]
			m2 := pair[1]
			m1.applyGravity(*m2)
			m2.applyGravity(*m1)
		}
		for _, moon := range moons {
			moon.move()
		}
	}
	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += moon.totalEnergy()
	}
	return totalEnergy
}

func calcStepsUntilRepeat(moons []*moon) int {
	moonPairs := getMoonPairs(moons)
	xPeriod := calcStepsUntilRepeatSinglePlane(moonPairs, moons,
		func(m *moon) int { return m.position.x },
		func(m *moon) int { return m.velocity.x },
		func(m *moon, v int) { m.velocity.x += v },
		func(m *moon, v int) { m.position.x += v })

	yPeriod := calcStepsUntilRepeatSinglePlane(moonPairs, moons,
		func(m *moon) int { return m.position.y },
		func(m *moon) int { return m.velocity.y },
		func(m *moon, v int) { m.velocity.y += v },
		func(m *moon, v int) { m.position.y += v })

	zPeriod := calcStepsUntilRepeatSinglePlane(moonPairs, moons,
		func(m *moon) int { return m.position.z },
		func(m *moon) int { return m.velocity.z },
		func(m *moon, v int) { m.velocity.z += v },
		func(m *moon, v int) { m.position.z += v })

	return leastCommonMultiple(xPeriod, yPeriod, zPeriod)
}

func calcStepsUntilRepeatSinglePlane(
	moonPairs [][]*moon,
	moons []*moon,
	getPos, getVel func(m *moon) int,
	setVel, setPos func(m *moon, v int)) int {

	moonStates := map[string]bool{}
	for i := 0; ; i++ {
		for _, pair := range moonPairs {
			m1 := pair[0]
			m2 := pair[1]
			m1.applyGravitySinglePlane(m2, getPos, setVel)
			m2.applyGravitySinglePlane(m1, getPos, setVel)
		}
		for _, moon := range moons {
			setPos(moon, getVel(moon))
		}
		var moonState strings.Builder
		moonState.Grow(8)
		for _, moon := range moons {
			_, _ = fmt.Fprintf(&moonState, "%d%d", getPos(moon),getVel(moon))
		}
		s := moonState.String()
		if moonStates[s] {
			return i
		}
		moonStates[s] = true
	}
}

func leastCommonMultiple(a, b, c int) int {
	aNext := a
	bNext := b
	cNext := c
	for {
		if aNext == bNext && bNext == cNext {
			return aNext
		}
		//increase the smallest one
		if aNext < bNext && aNext < cNext {
			aNext += a
		} else if bNext < cNext {
			bNext += b
		} else {
			cNext += c
		}
	}
}

func getMoonPairs(moons []*moon) [][]*moon {
	var moonPairs [][]*moon
	for first := 0; first < len(moons)-1; first++ {
		for second := first + 1; second < len(moons); second++ {
			moonPairs = append(moonPairs, []*moon{moons[first], moons[second]})
		}
	}
	return moonPairs
}
