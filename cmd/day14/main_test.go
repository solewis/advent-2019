package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseReactions(t *testing.T) {
	t.Run("sample 1", func(t *testing.T) {
		input := `10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`
		reactions := parseReactions(strings.Split(input, "\n"))
		expected := map[string]reaction{
			"FUEL": {requirements: map[string]int{"A": 7, "E": 1}, output: 1},
			"E":    {requirements: map[string]int{"A": 7, "D": 1}, output: 1},
			"D":    {requirements: map[string]int{"A": 7, "C": 1}, output: 1},
			"C":    {requirements: map[string]int{"A": 7, "B": 1}, output: 1},
			"B":    {requirements: map[string]int{"ORE": 1}, output: 1},
			"A":    {requirements: map[string]int{"ORE": 10}, output: 10},
		}
		if !reflect.DeepEqual(reactions, expected) {
			t.Errorf("Parse failed")
		}
	})
}

func TestCalcRequiredOre(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		reactions := map[string]reaction{
			"FUEL": {requirements: map[string]int{"A": 7, "B": 1}, output: 1},
			"B":    {requirements: map[string]int{"ORE": 1}, output: 1},
			"A":    {requirements: map[string]int{"ORE": 10}, output: 10},
		}
		result := calcRequiredOre(reactions, 1)
		expected := 11
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
	t.Run("uses store", func(t *testing.T) {
		reactions := map[string]reaction{
			"FUEL": {requirements: map[string]int{"A": 7, "C": 1}, output: 1},
			"B":    {requirements: map[string]int{"ORE": 1}, output: 1},
			"C":    {requirements: map[string]int{"A": 2, "B": 1}, output: 1},
			"A":    {requirements: map[string]int{"ORE": 10}, output: 10},
		}
		result := calcRequiredOre(reactions, 1)
		expected := 11
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
	t.Run("uses store order independent", func(t *testing.T) {
		reactions := map[string]reaction{
			//ensure c coming first still takes advantage of A's excess
			"FUEL": {requirements: map[string]int{"C": 1, "A": 7}, output: 1},
			"B":    {requirements: map[string]int{"ORE": 1}, output: 1},
			"C":    {requirements: map[string]int{"A": 2, "B": 1}, output: 1},
			"A":    {requirements: map[string]int{"ORE": 10}, output: 10},
		}
		result := calcRequiredOre(reactions, 1)
		expected := 11
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("sample 1", func(t *testing.T) {
		reactions := map[string]reaction{
			"FUEL": {requirements: map[string]int{"A": 7, "E": 1}, output: 1},
			"E":    {requirements: map[string]int{"A": 7, "D": 1}, output: 1},
			"D":    {requirements: map[string]int{"A": 7, "C": 1}, output: 1},
			"C":    {requirements: map[string]int{"A": 7, "B": 1}, output: 1},
			"B":    {requirements: map[string]int{"ORE": 1}, output: 1},
			"A":    {requirements: map[string]int{"ORE": 10}, output: 10},
		}
		result := calcRequiredOre(reactions, 1)
		expected := 31
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("sample 2", func(t *testing.T) {
		input := `9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL`
		result := calcRequiredOre(parseReactions(strings.Split(input, "\n")), 1)
		expected := 165
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("sample 3", func(t *testing.T) {
		input := `157 ORE => 5 D
165 ORE => 6 H
44 A, 5 B, 1 C, 29 D, 9 E, 48 F => 1 FUEL
12 F, 1 E, 8 G => 9 C
179 ORE => 7 G
177 ORE => 5 F
7 H, 7 G => 2 A
165 ORE => 2 E
3 H, 7 D, 5 F, 10 G => 8 B`
		result := calcRequiredOre(parseReactions(strings.Split(input, "\n")), 1)
		expected := 13312
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("sample 4", func(t *testing.T) {
		input := `2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF`
		result := calcRequiredOre(parseReactions(strings.Split(input, "\n")), 1)
		expected := 180697
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("sample 5", func(t *testing.T) {
		input := `171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`
		result := calcRequiredOre(parseReactions(strings.Split(input, "\n")), 1)
		expected := 2210736
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
}

func TestCalcMaxFuel(t *testing.T) {
	t.Run("sample 1", func(t *testing.T) {
		input := `157 ORE => 5 D
165 ORE => 6 H
44 A, 5 B, 1 C, 29 D, 9 E, 48 F => 1 FUEL
12 F, 1 E, 8 G => 9 C
179 ORE => 7 G
177 ORE => 5 F
7 H, 7 G => 2 A
165 ORE => 2 E
3 H, 7 D, 5 F, 10 G => 8 B`
		result := calcMaxFuel(parseReactions(strings.Split(input, "\n")), 13312)
		expected := 82892753
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("sample 2", func(t *testing.T) {
		input := `2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF`
		result := calcMaxFuel(parseReactions(strings.Split(input, "\n")), 180697)
		expected := 5586022
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})

	t.Run("sample 3", func(t *testing.T) {
		input := `171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`
		result := calcMaxFuel(parseReactions(strings.Split(input, "\n")), 2210736)
		expected := 460664
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
}
