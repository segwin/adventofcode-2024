package day18

import (
	"errors"
	"fmt"
	"slices"

	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

var (
	errUnsolved = errors.New("failed to solve maze")
)

type Solution struct {
	FallingBytes []map2d.Position
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 18:\n")

	const mapSize = 71 // from problem statement
	memory := NewEmptyLayout(mapSize, mapSize)
	start := map2d.Position{X: 0, Y: 0}
	end := map2d.Position{X: mapSize - 1, Y: mapSize - 1}

	fallenBytes := 1024
	memory = DropBytes(memory, s.FallingBytes[:1024]...)

	// part 1
	bestPath := Solve(memory, start, end)
	if bestPath == nil {
		return errUnsolved
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Shortest path: %d\n", Distance(bestPath))

	// part 2
	firstBlocking := FirstBlockingByte(memory, start, end, s.FallingBytes[fallenBytes:])

	fmt.Print("  PART 2:\n")
	fmt.Printf("    First blocking byte's coordinates: %d,%d\n", firstBlocking.X, firstBlocking.Y)

	return nil
}

// DropBytes adds the given falling bytes to the layout, causing their positions to become corrupted.
func DropBytes(memory Layout, positions ...map2d.Position) Layout {
	for _, pos := range positions {
		memory = memory.With(pos, Corrupted)
	}
	return memory
}

type Path []map2d.Position

// Solve finds the shortest path(s) through the maze. More than one path may be returned if multiple
// paths result in the same optimal distance.
//
// The implementation is mostly copied from day 16's approach, but without the scoring constraint.
func Solve(memory Layout, start, end map2d.Position) Path {
	// use Dijkstra's algorithm like in day 16
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Algorithm

	// 1: construct the set of unvisted nodes
	// 2: assign each node a distance from start (initially infinite for non-start nodes)
	unvisited := unvisitedNodes(memory, start)

	bestPathByPos := map[map2d.Position]Path{}
	for len(unvisited) > 0 {
		// 3: select the lowest-cost node as current
		var cur map2d.Position
		var curPath Path
		for node, path := range unvisited {
			if path == nil {
				continue // unresolved node
			}
			if curPath == nil || len(path) < len(curPath) {
				curPath = path
				cur = node
			}
		}

		if curPath == nil {
			break // go to 6: only unreachable nodes remain
		}

		// 4: compute costs for all unvisited neighbouring nodes
		for _, nextDir := range map2d.CardinalDirections() {
			nextPos := cur.Move(nextDir, 1)
			if tile, ok := memory.Get(nextPos); !ok || tile == Corrupted {
				continue // not a valid path candidate
			}
			if _, isUnvisited := unvisited[nextPos]; !isUnvisited {
				continue // already visited
			}

			// ok: next tile is a valid path candidate, see if it's better than anything we found before
			nextPath := append(slices.Clone(curPath), nextPos)
			unvisited[nextPos] = updatedBestPath(unvisited[nextPos], nextPath)
		}

		// 5: remove current node from the unvisited set
		delete(unvisited, cur)
		bestPathByPos[cur] = updatedBestPath(bestPathByPos[cur], curPath)
	}

	// 6: end node is now assigned the lowest possible score (shortest route)
	return bestPathByPos[end]
}

// Distance returns the number of steps along the given path.
func Distance(path Path) int {
	return len(path) - 1
}

// FirstBlockingByte returns the coordinates of the first falling byte that causes the memory layout
// to become inescapable (end walled off from start).
func FirstBlockingByte(memory Layout, start, end map2d.Position, fallingBytes []map2d.Position) map2d.Position {
	// cut the search space in half on each attempt
	// examples:
	//  - blocking byte is 1st of 32: 16 -> 8 -> 4 -> 2 -> 1 -> 0
	//  - blocking byte is 5th of 32: 16 -> 8 -> 4 -> 6 -> 5
	//  - blocking byte is 8th of 32: 16 -> 8 -> 4 -> 6 -> 7 -> 8
	//  - blocking byte is 21st of 32: 16 -> 24 -> 20 -> 22 -> 21
	//  - blocking byte is 24th of 32: 16 -> 24 -> 20 -> 22 -> 23 -> 24
	//  - blocking byte is 32nd of 32: 16 -> 24 -> 28 -> 30 -> 31

	lastNonBlocking := 0                   // first byte should not be blocking
	firstBlocking := len(fallingBytes) - 1 // must be blocked by the time the last byte falls

	for {
		dropCount := lastNonBlocking + max(1, (firstBlocking-lastNonBlocking)/2)
		bestPath := Solve(DropBytes(memory, fallingBytes[:dropCount]...), start, end)
		if bestPath == nil {
			firstBlocking = dropCount - 1
		} else {
			lastNonBlocking = dropCount
		}

		if firstBlocking-lastNonBlocking == 1 {
			return fallingBytes[firstBlocking] // found first blocking byte
		}
	}
}

// unvisitedNodes builds and returns the set of all unvisited nodes in the maze, including the start
// and end tiles. At first, only the starting position has a known path (distance = 0).
func unvisitedNodes(memory Layout, start map2d.Position) map[map2d.Position]Path {
	unvisited := map[map2d.Position]Path{}
	for i, row := range memory {
		for j, tile := range row {
			if tile == Empty {
				pos := map2d.PositionFromIndex(i, j)
				unvisited[pos] = nil
			}
		}
	}

	// starting position always has distance = 0
	unvisited[start] = Path{start}

	return unvisited
}

// updatedBestPath returns the new shortest path between the current & new one.
func updatedBestPath(bestPath, newPath Path) Path {
	if bestPath == nil {
		return newPath // anything is better than nothing
	}
	if newPath == nil {
		return bestPath // no new path found
	}

	if len(newPath) < len(bestPath) {
		return newPath // better: replace best path
	}
	return bestPath // worse or equivalent: keep the current best
}
