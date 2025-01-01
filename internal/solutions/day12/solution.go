package day12

import (
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Solution struct {
	Garden Garden
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 12:\n")

	regions := GetRegions(s.Garden)

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Total fence price: %d\n", TotalFencePrice(regions))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Total discounted fence price: %d\n", TotalDiscountedFencePrice(regions, s.Garden))

	return nil
}

// GetRegions aggregates all adjacent plots of the same kind in the garden into regions.
func GetRegions(garden Garden) []Region {
	regions := make([]Region, 0, len(garden)) // worst case: # regions = # plots

	consumedPositions := map[map2d.Position]struct{}{} // positions already assigned to a region
	for i := range garden {
		for j, kind := range garden[i] {
			pos := map2d.PositionFromIndex(i, j)
			if _, consumed := consumedPositions[pos]; consumed {
				continue // already covered this plot
			}

			positions, perimeter := findRegionAround(pos, kind, garden, consumedPositions)
			regions = append(regions, Region{
				Kind:      kind,
				Positions: positions,
				Area:      len(positions),
				Perimeter: perimeter,
			})
		}
	}

	return regions
}

func findRegionAround(
	cur map2d.Position,
	kind byte,
	in Garden,
	consumedPositions map[map2d.Position]struct{},
) (positions map[map2d.Position]struct{}, perimeter int) {
	// store current position
	positions = map[map2d.Position]struct{}{cur: {}}
	consumedPositions[cur] = struct{}{}

	// find adjacent plots of the same kind
	for _, dir := range map2d.CardinalDirections() {
		next := cur.Move(dir, 1)
		nextPlot, ok := in.Get(next)
		if !ok {
			perimeter++
			continue // reached garden edge
		}
		if nextPlot != kind {
			perimeter++
			continue // edge within garden
		}

		if _, consumed := consumedPositions[next]; consumed {
			continue // same kind, but already used
		}

		// recurse along this path, but don't come back to this block
		nextPositions, nextPerimeter := findRegionAround(next, kind, in, consumedPositions)
		for pos := range nextPositions {
			positions[pos] = struct{}{}
		}
		perimeter += nextPerimeter
	}

	return positions, perimeter
}

// TotalFencePrice returns the total cost of fencing for these regions without any discounts.
func TotalFencePrice(regions []Region) int {
	sum := 0
	for _, r := range regions {
		sum += fencePrice(r)
	}
	return sum
}

func fencePrice(r Region) int {
	return r.Area * r.Perimeter
}

func TotalDiscountedFencePrice(regions []Region, garden Garden) int {
	sum := 0
	for _, r := range regions {
		sum += discountedFencePrice(r, garden)
	}
	return sum
}

func discountedFencePrice(r Region, garden Garden) int {
	return r.Area * r.Sides(garden)
}
