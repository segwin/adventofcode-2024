package day20

import (
	"fmt"
	"math"
	"slices"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Solution struct {
	Racetrack map2d.Map[Tile]
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 20:\n")

	start, ok := s.Racetrack.Find(Start)
	if !ok {
		return fmt.Errorf("%w: start tile not found", parsing.ErrInvalidData)
	}
	end, ok := s.Racetrack.Find(End)
	if !ok {
		return fmt.Errorf("%w: end tile not found", parsing.ErrInvalidData)
	}

	normalRoute := Navigate(s.Racetrack, start, end)

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Cheats that save at least 100ps: %d\n", len(FindCheats(s.Racetrack, normalRoute, 100)))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    TODO\n")

	return nil
}

// Navigate the racetrack, returning all steps on the regular path without any cheats.
// The path does not include the starting tile.
func Navigate(racetrack map2d.Map[Tile], start, end map2d.Position) (path []map2d.Position) {
	cur := start
	for cur != end {
		for _, dir := range map2d.CardinalDirections() {
			next := cur.Move(dir, 1)
			if len(path) != 0 && next == path[len(path)-1] {
				continue // don't backtrack
			}
			if nextTile, ok := racetrack.Get(next); !ok || nextTile == Wall {
				continue // stay on the track
			}

			path = append(path, cur)
			cur = next
			break
		}
	}

	path = append(path, end) // epilogue: add end position
	return path
}

// FindCheats returns all cheats along the race path, mapped to their time savings in picoseconds.
// A position is a cheat candidate if it holds a wall that lies between two track tiles.
func FindCheats(racetrack map2d.Map[Tile], normalRoute []map2d.Position, minSavings int) (cheats map[Cheat]int) {
	cheats = map[Cheat]int{}
	for _, trackPos := range normalRoute {
		for _, dir := range map2d.CardinalDirections() {
			start := trackPos.Move(dir, 1)
			if tile, ok := racetrack.Get(start); !ok || tile != Wall {
				continue // cheats are only helpful for blipping through walls
			}

			end := start.Move(dir, 1)
			if tile, ok := racetrack.Get(end); !ok || tile == Wall {
				continue // cheat doesn't lead to the racetrack
			}

			cheat := Cheat{Start: start, End: end}
			if _, exists := cheats[cheat]; exists {
				continue // already identified this cheat earlier in the route
			}
			if savings := timeSavings(cheat, normalRoute); savings >= minSavings {
				cheats[cheat] = savings
			}
		}
	}

	return cheats
}

func timeSavings(cheat Cheat, normalRoute []map2d.Position) int {
	for startIdx, pos := range normalRoute {
		if !cheat.Start.AdjacentTo(pos) {
			continue // cheat doesn't apply here
		}
		endIdx := slices.Index(normalRoute, cheat.End)
		return endIdx - startIdx - 2
	}
	return math.MinInt // -infinity: invalid cheat
}
