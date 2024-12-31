package day8

import (
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Solution struct {
	CityMap CityMap
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 8:\n")

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Unique antinode locations: %d\n", UniqueAntinodeLocations(s.CityMap, false))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Unique antinode locations, with harmonics: %d\n", UniqueAntinodeLocations(s.CityMap, true))

	return nil
}

func UniqueAntinodeLocations(cityMap CityMap, includeHarmonics bool) int {
	uniqueLocations := map[map2d.Position]struct{}{}
	for _, antinode := range findAllAntinodes(cityMap, includeHarmonics) {
		uniqueLocations[antinode] = struct{}{}
	}
	return len(uniqueLocations)
}

func findAllAntinodes(cityMap CityMap, includeHarmonics bool) []map2d.Position {
	// catalog all antennae on the map by type
	antennae := map[Tile][]map2d.Position{}
	for i, row := range cityMap {
		for j, tile := range row {
			if tile == Empty {
				continue // nothing to do
			}
			antennae[tile] = append(antennae[tile], map2d.PositionFromIndex(i, j))
		}
	}

	// for each type, find each antenna pair's antinodes
	var antinodes []map2d.Position
	for _, positions := range antennae {
		if len(positions) < 2 {
			continue // antinodes require at least 1 pair
		}

		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				// find all antinodes along the antenna pair's path, in each direction
				p1, p2 := positions[i], positions[j]
				antinodes = append(antinodes, findAntinodesAfter(cityMap, p1, p2, includeHarmonics)...)
				antinodes = append(antinodes, findAntinodesAfter(cityMap, p2, p1, includeHarmonics)...)
			}
		}
	}

	return antinodes
}

func findAntinodesAfter(cityMap CityMap, p1, p2 map2d.Position, includeHarmonics bool) (antinodes []map2d.Position) {
	distance := p2.Sub(p1)

	if includeHarmonics {
		antinodes = []map2d.Position{p1, p2} // antennae's positions are included in harmonics mode
	}

	candidate := p1.Sub(distance)
	for {
		if _, ok := cityMap.Get(candidate); !ok {
			return antinodes // reached end of map
		}

		antinodes = append(antinodes, candidate)
		if !includeHarmonics {
			return antinodes // harmonics disabled, don't calculate other antinodes along the path
		}

		candidate = candidate.Sub(distance)
	}
}
