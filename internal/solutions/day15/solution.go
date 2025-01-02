package day15

import (
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Solution struct {
	Layout       Layout
	Instructions []map2d.Direction
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 15:\n")

	// part 1
	finalLayout, err := Resolve(s.Layout, s.Instructions...)
	if err != nil {
		return fmt.Errorf("resolving instructions in regular layout: %w", err)
	}

	gpsSum := 0
	for _, box := range finalLayout.FindAll(Box) {
		gpsSum += 100*box.Y + box.X
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Box GPS sum: %d\n", gpsSum)

	// part 2
	finalLayout, err = Resolve(WidenWarehouse(s.Layout), s.Instructions...)
	if err != nil {
		return fmt.Errorf("resolving instructions in wide layout: %w", err)
	}

	gpsSum = 0
	for _, box := range finalLayout.FindAll(BoxL) {
		gpsSum += 100*box.Y + box.X
	}

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Box GPS sum in wide warehouse: %d\n", gpsSum)

	return nil
}

// Resolve all instructions for the robot, returning the final updated layout.
// Returns an error if the robot can't be found in layout.
func Resolve(layout Layout, instructions ...map2d.Direction) (Layout, error) {
	robot, ok := layout.Find(Robot)
	if !ok {
		return Layout{}, fmt.Errorf("%w: robot not found", parsing.ErrInvalidData)
	}

	for _, instruction := range instructions {
		layout, robot = resolveOne(layout, robot, instruction)
	}
	return layout, nil
}

func resolveOne(layout Layout, robot map2d.Position, instruction map2d.Direction) (Layout, map2d.Position) {
	// find all boxes in the way in order to move them too
	movementFront := map[map2d.Position]struct{}{robot: {}} // position of all objects at the front of computed movement
	tilesToMove := map[map2d.Position]Tile{robot: Robot}
	for len(movementFront) != 0 {
		newMovementFront := map[map2d.Position]struct{}{}
		for frontPos := range movementFront {
			nextPos := frontPos.Move(instruction, 1)
			nextTile, _ := layout.Get(nextPos) // can never go out of bounds since there's an impassable wall

			// handle other half of box in extra-wide map
			switch nextTile {
			case Wall:
				return layout, robot // stop immediately: something is hitting a wall
			case Empty:
				continue // ok: nothing but air ahead
			}

			// tile ahead is a box: add it to the next iteration's movement front
			newMovementFront[nextPos] = struct{}{}
			tilesToMove[nextPos] = nextTile

			if instruction == map2d.East() || instruction == map2d.West() {
				continue // stop here: wide (L/R) boxes only require joining in vertical movement
			}

			// moving vertically: ensure paired L/R boxes are moved together
			switch nextTile {
			case BoxL:
				joinedPos := nextPos.Move(map2d.East(), 1)
				tilesToMove[joinedPos] = BoxR
				newMovementFront[joinedPos] = struct{}{}
			case BoxR:
				joinedPos := nextPos.Move(map2d.West(), 1)
				tilesToMove[joinedPos] = BoxL
				newMovementFront[joinedPos] = struct{}{}
			}
		}
		movementFront = newMovementFront
	}

	// ok: move all tiles to their new positions
	for pos := range tilesToMove {
		layout.Map = layout.With(pos, Empty)
	}
	for pos, tile := range tilesToMove {
		layout.Map = layout.With(pos.Move(instruction, 1), tile)
	}

	return layout, robot.Move(instruction, 1)
}

// WidenWarehouse returns the "wide" version of this layout, with twice as many columns and tiles
// adjusted using the following rules:
//
//   - # => ##
//   - O => []
//   - . => ..
//   - @ => @.
func WidenWarehouse(layout Layout) (wideLayout Layout) {
	wideLayout.Map = make([][]Tile, len(layout.Map)) // same number of rows
	for i, row := range layout.Map {
		wideLayout.Map[i] = make([]Tile, 2*len(row)) // twice as many columns
		for j, tile := range row {
			switch tile {
			case Robot: // @ => @.
				wideLayout.Map[i][2*j+0] = Robot
				wideLayout.Map[i][2*j+1] = Empty
			case Box: // O => []
				wideLayout.Map[i][2*j+0] = BoxL
				wideLayout.Map[i][2*j+1] = BoxR
			default: // everything else: tile is doubled up
				wideLayout.Map[i][2*j+0] = tile
				wideLayout.Map[i][2*j+1] = tile
			}
		}
	}
	return wideLayout
}
