package day16

import (
	"errors"
	"fmt"
	"slices"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

var (
	errUnsolved = errors.New("failed to solve maze")
)

type Solution struct {
	Maze Maze
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 16:\n")

	start, ok := s.Maze.Find(Start)
	if !ok {
		return fmt.Errorf("%w: start position not found", parsing.ErrInvalidData)
	}
	end, ok := s.Maze.Find(End)
	if !ok {
		return fmt.Errorf("%w: end position not found", parsing.ErrInvalidData)
	}

	bestPaths := Solve(s.Maze, start, end)
	if len(bestPaths) == 0 {
		return errUnsolved
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Lowest possible score: %d\n", bestPaths[0].Score)

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Number of seats: %d\n", CountSeats(bestPaths))

	return nil
}

// Solve finds the optimal path(s) through the maze. More than one path may be returned if multiple
// paths result in the same optimal score.
func Solve(maze Maze, start, end map2d.Position) []Path {
	// use Dijkstra's algorithm since this is essentially a weighted graph (turning costs more than moving forward)
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Algorithm

	// 1: construct the set of unvisted nodes
	// 2: assign each node a distance from start (initially infinite for non-start nodes)
	unvisited := unvisitedNodes(maze, start)

	bestPathsByPos := map[map2d.Position][]Path{} // multiple paths may have the same score
	for len(unvisited) > 0 {
		// 3: select the lowest-cost node as current
		var cur pathNode
		var curScore int
		var curPaths []Path
		for node, equivPaths := range unvisited {
			if len(equivPaths) == 0 {
				continue // unresolved node
			}
			if len(curPaths) == 0 || equivPaths[0].Score < curScore {
				// better path: replace everything
				curPaths = equivPaths
				curScore = equivPaths[0].Score
				cur = node
			}
		}

		if len(curPaths) == 0 {
			break // go to 6: only unreachable nodes remain
		}

		// 4: compute costs for all unvisited neighbouring nodes
		for _, nextDir := range map2d.CardinalDirections() {
			nextPos := cur.Position.Move(nextDir, 1)
			if tile, ok := maze.Get(nextPos); !ok || tile == Wall {
				continue // not a valid path candidate
			}

			nextNode := pathNode{Position: nextPos, Direction: nextDir}
			if _, isUnvisited := unvisited[nextNode]; !isUnvisited {
				continue // already visited
			}

			// ok: next tile is a valid path candidate, see if it's better than anything we found before
			nextScore := curScore + 1 + 1000*(nextDir.Angle(cur.Direction)/90) // moving forward costs 1, turning 90 degrees costs 1000
			nextPaths := make([]Path, len(curPaths))
			for i, curPath := range curPaths {
				nextPaths[i] = Path{
					Score: nextScore,
					Steps: append(slices.Clone(curPath.Steps), nextPos),
				}
			}

			unvisited[nextNode] = updatedBestPaths(unvisited[nextNode], nextPaths)
		}

		// 5: remove current node from the unvisited set
		delete(unvisited, cur)
		bestPathsByPos[cur.Position] = updatedBestPaths(bestPathsByPos[cur.Position], curPaths)
	}

	// 6: end node is now assigned the lowest possible score (shortest route)
	return bestPathsByPos[end]
}

// CountSeats returns the number of unique seats along the given paths through the maze, AKA the number
// of unique positions along all paths.
func CountSeats(paths []Path) int {
	uniquePositions := map[map2d.Position]struct{}{}
	for _, path := range paths {
		for _, pos := range path.Steps {
			uniquePositions[pos] = struct{}{}
		}
	}
	return len(uniquePositions)
}

// unvisitedNodes builds and returns the set of all unvisited nodes in the maze, including the start
// and end tiles. At first, only the starting position has a known path (distance = 0).
func unvisitedNodes(maze Maze, start map2d.Position) map[pathNode][]Path {
	unvisited := map[pathNode][]Path{
		// start facing east
		{Position: start, Direction: map2d.East()}: {{Score: 0, Steps: []map2d.Position{start}}},
	}
	for i, row := range maze {
		for j, tile := range row {
			if tile == Empty || tile == End {
				// 2: all other nodes have infinite (unknown) cost
				pos := map2d.PositionFromIndex(i, j)
				for _, dir := range map2d.CardinalDirections() {
					unvisited[pathNode{Position: pos, Direction: dir}] = nil // no paths known
				}
			}
		}
	}
	return unvisited
}

// updatedBestPaths returns the new set of "best" (lowest-scored) paths considering newPaths.
func updatedBestPaths(bestPaths, newPaths []Path) []Path {
	if len(bestPaths) == 0 {
		return newPaths // anything is better than nothing
	}
	if len(newPaths) == 0 {
		return bestPaths // no new paths found
	}

	if bestScore, newScore := bestPaths[0].Score, newPaths[0].Score; newScore < bestScore {
		return newPaths // better score: replace best paths
	} else if newScore > bestScore {
		return bestPaths // worse score: current is already best
	}
	return append(slices.Clone(bestPaths), newPaths...) // same score: add new paths to current best
}

type pathNode struct {
	// Position of this node in the path, i.e. the tip of the path.
	Position map2d.Position
	// Direction the reindeer is facing as of this node.
	Direction map2d.Direction
}

type Path struct {
	// Score of the path up to this node.
	Score int
	// Steps holds each position that leads up to this node, excluding itself.
	Steps []map2d.Position
}
