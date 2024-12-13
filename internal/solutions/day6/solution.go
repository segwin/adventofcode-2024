package day6

import (
	"errors"
	"fmt"
	"sync"

	"github.com/segwin/adventofcode-2024/internal/parsing"
)

var (
	errLoopFound = errors.New("guard is in a loop")
)

type Solution struct {
	FloorMap FloorMap
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 6:\n")

	uniquePositions, states, err := CountGuardPositions(s.FloorMap)
	if err != nil {
		return fmt.Errorf("calculating guard positions: %w", err)
	}
	fmt.Print("  PART 1:\n")
	fmt.Printf("    Guard visited tiles: %d\n", uniquePositions)

	loopPositions, err := CountLoopPositions(s.FloorMap, states)
	if err != nil {
		return fmt.Errorf("calculating potential loops: %w", err)
	}
	fmt.Print("  PART 2:\n")
	fmt.Printf("    Loop candidates: %d\n", loopPositions)

	return nil
}

// CountGuardPositions returns the number of distinct positions the guard occupies on the map before
// leaving the room. It also returns the updated map showing the path the guard took.
func CountGuardPositions(floorMap FloorMap) (uniquePositions int, states []GuardState, err error) {
	initialStates := findGuardStates(floorMap)
	if len(initialStates) != 1 {
		return 0, nil, fmt.Errorf("%w: expected 1 guard state, got %d", parsing.ErrInvalidData, len(states))
	}

	guard := &initialStates[0]
	for guard != nil {
		for _, previous := range states {
			if guard.Position == previous.Position && guard.Direction == previous.Direction {
				return 0, nil, errLoopFound // guard is stuck in a loop
			}
		}
		states = append(states, *guard)
		guard, floorMap = guard.AdvanceOne(floorMap)
	}

	for _, row := range floorMap {
		for _, tile := range row {
			if tile.IsGuard() {
				uniquePositions++
			}
		}
	}
	return uniquePositions, states, nil
}

// CountLoopPositions returns the number of positions on the map where adding an obstacle would cause
// the guard to enter a loop.
//
// Loop candidates:
//   - are somewhere on the guard's normal path
//   - cause the path to form a closed loop with existing obstacles
func CountLoopPositions(floorMap FloorMap, originalStates []GuardState) (int, error) {
	startingPosition := originalStates[0].Position // starting position is protected

	// produce loop positions by propagating each state after inserting the obstacle
	var wg sync.WaitGroup
	loopPositions := make(chan Position)
	errs := make(chan error, 1)

	for i := 0; i < len(originalStates)-1; i++ {
		cur, next := originalStates[i], originalStates[i+1]

		wg.Add(1)
		go func() {
			defer wg.Done()

			if next.Position == startingPosition {
				return // can't block the guard's starting position
			}

			foundLoop, err := obstacleAheadCausesLoop(cur, next, floorMap)
			if err != nil {
				errs <- err
				return
			}
			if foundLoop {
				loopPositions <- next.Position
			}
		}()
	}

	// once all producers are done, close the channel to end collection
	go func() {
		wg.Wait()
		close(loopPositions)
	}()

	// collect loop positions as they're produced
	uniqueLoopPositions := map[Position]struct{}{} // deduplicate obstacles that generate >1 loop
	var err error

	for {
		select {
		case pos, ok := <-loopPositions:
			if !ok {
				return len(uniqueLoopPositions), err
			}
			uniqueLoopPositions[pos] = struct{}{}

		case producerErr := <-errs:
			err = errors.Join(err, producerErr) // continue to collect all individual errors
		}
	}
}

// obstacleAheadCausesLoop returns true if placing a single obstacle in front of the guard causes
// them to enter a loop.
func obstacleAheadCausesLoop(cur, next GuardState, floorMap FloorMap) (bool, error) {
	if next.Position == cur.Position {
		return false, nil // guard is not moving between these states
	}

	// try putting an obstacle at the next position
	alteredMap := floorMap.
		WithTile(next.Position, Obstacle)

	_, _, err := CountGuardPositions(alteredMap)
	if errors.Is(err, errLoopFound) {
		return true, nil
	}
	if err != nil {
		return false, fmt.Errorf("non-loop error: %w", err)
	}
	return false, nil
}

func findGuardStates(floorMap FloorMap) (states []GuardState) {
	for i, row := range floorMap {
		for j, tile := range row {
			if !tile.IsGuard() {
				continue // not a guard
			}
			states = append(states, GuardState{
				Position:  Position{X: j, Y: i},
				Direction: tile.Direction(),
			})
		}
	}
	return states
}
