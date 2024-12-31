package day10

import (
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Solution struct {
	Terrain TopographicMap
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 10:\n")

	trailheads := GetTrailheads(s.Terrain)

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Total score: %d\n", AddScores(trailheads))

	return nil
}

func GetTrailheads(terrain TopographicMap) []Trailhead {
	var trailheads []Trailhead
	for i, row := range terrain {
		for j, cell := range row {
			if cell != MinHeight {
				continue // trailheads are always at height = 0
			}

			position := map2d.PositionFromIndex(i, j)
			root := climbAround(terrain, position)

			trailheads = append(trailheads, Trailhead{
				Position: position,
				Score:    root.CountTrails(),
			})
		}
	}

	return trailheads
}

func AddScores(trailheads []Trailhead) int {
	sum := 0
	for _, th := range trailheads {
		sum += th.Score
	}
	return sum
}

func climbAround(terrain TopographicMap, curPos map2d.Position) TrailStep {
	cur, ok := terrain.Get(curPos)
	if !ok {
		return TrailStep{Position: curPos, Next: nil} // invalid position, can't climb from here
	}

	root := TrailStep{Position: curPos, Height: cur}
	for _, dir := range []map2d.Direction{map2d.North(), map2d.East(), map2d.South(), map2d.West()} {
		nextPos := curPos.Move(dir, 1)
		next, ok := terrain.Get(nextPos)
		if !ok || next-cur != 1 {
			continue // can't climb here
		}
		root.Next = append(root.Next, climbAround(terrain, nextPos))
	}

	return root
}
