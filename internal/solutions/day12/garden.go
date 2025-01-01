package day12

import (
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Garden = map2d.Map[byte]

type Region struct {
	Kind      byte
	Positions map[map2d.Position]struct{}
	Area      int
	Perimeter int
}

func (r *Region) Sides(garden Garden) int {
	horizontal := r.sidesInDirection(map2d.East(), map2d.South(), garden)
	vertical := r.sidesInDirection(map2d.South(), map2d.East(), garden)
	return horizontal + vertical
}

func (r *Region) sidesInDirection(dir map2d.Direction, crossDir map2d.Direction, garden Garden) int {
	// store edges we find:
	//   - key: position component aligned with dir
	//   - val: list of position components perpendicular to dir, along the key's axis
	enteringEdges := map[int][]int{}
	exitingEdges := map[int][]int{}

	// find edges along the given direction, navigating through the map
	for crossPos := (map2d.Position{X: 0, Y: 0}); garden.Contains(crossPos); crossPos = crossPos.Move(crossDir, 1) { // cross direction
		inRegion := false

		cur := crossPos
		for ; garden.Contains(cur); cur = cur.Move(dir, 1) { // edge direction
			// is cur in the region?
			if _, ok := r.Positions[cur]; ok {
				if !inRegion { // found transition point
					inRegion = true
					along, across := cur.ProjectOnto(dir), cur.ProjectOnto(crossDir)
					enteringEdges[along] = append(enteringEdges[along], across)
				}

				continue
			}

			// not in the region
			if inRegion { // found transition point
				inRegion = false
				along, across := cur.ProjectOnto(dir), cur.ProjectOnto(crossDir)
				exitingEdges[along] = append(exitingEdges[along], across)
				continue
			}
		}

		if inRegion {
			// we left the region upon exiting the garden
			along, across := cur.ProjectOnto(dir), cur.ProjectOnto(crossDir)
			exitingEdges[along] = append(exitingEdges[along], across)
		}
	}

	// aggregate contiguous edges into sides
	sides := 0
	for _, edges := range []map[int][]int{enteringEdges, exitingEdges} {
		for _, edgeSections := range edges {
			sides++ // prologue: cover 1st y value
			for i := 1; i < len(edgeSections); i++ {
				if edgeSections[i]-edgeSections[i-1] > 1 { // non-contiguous => separate sides
					sides++
				}
			}
		}
	}

	return sides
}
