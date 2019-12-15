package main

import (
	"fmt"
	"github.com/solewis/advent-2019/internal/parse"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := parse.Lines("cmd/day14/input.txt")
	reactions := parseReactions(input)
	orePerFuel := calcRequiredOre(reactions, 1)
	fmt.Printf("Part 1: %d\n", orePerFuel)
	fmt.Printf("Part 2: %d\n", calcMaxFuel(reactions, orePerFuel))
}

type reaction struct {
	requirements map[string]int
	output       int
}

func calcMaxFuel(reactions map[string]reaction, orePerFuel int) int {
	// the minimum amount of fuel it could produce is the amount of ore divided by the orePerFuel
	// so we start there
	// but that misses out on utilizing the leftovers from the production of 1 FUEL in the next

	// by brute force, see if you can produce more and more FUEL with your trillion ORE, until you reach the max
	// do an exponential increment of the amount of FUEL tested each time for performance reasons
	trillion := 1000000000000
	i := trillion / orePerFuel // minimum possible FUEL
	increment := 1000000
	for {
		for {
			ore := calcRequiredOre(reactions, i)
			if ore > trillion {
				if increment == 1 {
					return i - 1
				}
				i-=increment // amount of ORE at the current FUEL level was more than a trillion, back off to previous test and lower the increment value
				break
			}
			i+=increment
		}
		increment /= 2 // exponential backoff, which will eventually get to 1
	}
}

func calcRequiredOre(reactions map[string]reaction, fuelNeeded int) int {
	var chemicals []string
	for c := range reactions {
		chemicals = append(chemicals, c)
	}
	chemicals = append(chemicals, "ORE")

	// Order chemicals based on depth to ORE. E.g ORE at the back, then the things that ORE makes in front of that, etc.
	sort.Slice(chemicals, func(i, j int) bool {
		return calcDepth(reactions, chemicals[i]) > calcDepth(reactions, chemicals[j])
	})

	amountsNeeded := map[string]int{"FUEL": fuelNeeded}
	for ic, c := range chemicals {
		// For each chemical, loop through all the preceding chemicals and figure out how much is needed to produce
		// the amount of that preceding chemical is needed. E.g. for FUEL we need 1 always. For the next chemical
		// determine how much is needed to satisfy 1 FUEL, then the next chemical decides how much it needs for 1 FUEL
		// plus n number of the second chemical, and so on.
		for ip := 0; ip < ic; ip++ {
			p := chemicals[ip]
			amountPNeeded := amountsNeeded[p]
			pReaction := reactions[p]
			for requirement, qty := range pReaction.requirements {
				if requirement == c {
					//record how much (c) is needed to fulfill the required amount of (p)
					amountCNeeded := int(math.Ceil(float64(amountPNeeded)/float64(pReaction.output))) * qty
					amountsNeeded[c] += amountCNeeded
				}
			}
		}
	}
	return amountsNeeded["ORE"]
}

func calcDepth(reactions map[string]reaction, chemical string) int {
	if _, exists := reactions[chemical]; !exists {
		return 0
	}
	reaction := reactions[chemical]
	maxDepth := 0
	for child := range reaction.requirements {
		depth := calcDepth(reactions, child)
		if depth > maxDepth {
			maxDepth = depth
		}
	}
	return 1 + maxDepth
}

func parseReactions(input []string) map[string]reaction {
	reactions := map[string]reaction{}
	re := regexp.MustCompile(`(\d+) ([A-Z]+)`)
	for _, l := range input {
		var reaction reaction
		parts := strings.Split(l, " => ")
		reqs := strings.Split(parts[0], ", ")
		requirements := map[string]int{}
		for _, req := range reqs {
			matches := re.FindStringSubmatch(req)
			ct, _ := strconv.Atoi(matches[1])
			requirements[matches[2]] = ct
		}
		reaction.requirements = requirements
		matches := re.FindStringSubmatch(parts[1])
		output, _ := strconv.Atoi(matches[1])
		reaction.output = output
		reactions[matches[2]] = reaction
	}
	return reactions
}
