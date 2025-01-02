package day16

import (
	"fmt"
	"math"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
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

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Lowest possible score: %d\n", Solve(s.Maze, start, end))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    TODO\n")

	return nil
}

// Solve finds the optimal path through the maze and returns its score.
func Solve(maze Maze, start, end map2d.Position) int {
	// use Dijkstra's algorithm since this is essentially a weighted graph (turning costs more than moving forward)
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Algorithm

	// 1: construct the set of unvisted nodes
	// 2: assign each node a distance from start (initially infinite for non-start nodes)
	unvisited := unvisitedNodes(maze, start)

	bestScoreByPos := map[map2d.Position]int{}
	for len(unvisited) > 0 {
		// 3: select the lowest-cost node as current
		var cur pathNode
		curScore := math.MaxInt
		for node, cost := range unvisited {
			if cost < curScore {
				curScore = cost
				cur = node
			}
		}

		if curScore == math.MaxInt {
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

			// ok: next tile is a valid path candidate, add it to the search stack
			nextScore := curScore + 1 + 1000*(nextDir.Angle(cur.Direction)/90) // moving forward costs 1, turning 90 degrees costs 1000
			if unvisited[nextNode] > nextScore {
				unvisited[nextNode] = nextScore
			}
		}

		// 5: remove current node from the unvisited set
		delete(unvisited, cur)
		if bestScore, ok := bestScoreByPos[cur.Position]; !ok || curScore < bestScore {
			bestScoreByPos[cur.Position] = curScore
		}
	}

	// 6: end node is now assigned the lowest possible score (shortest route)
	return bestScoreByPos[end]
}

func unvisitedNodes(maze Maze, start map2d.Position) map[pathNode]int {
	unvisited := map[pathNode]int{
		{Position: start, Direction: map2d.East()}:  0,    // no turn: cost = 0
		{Position: start, Direction: map2d.North()}: 1000, // 90deg turn
		{Position: start, Direction: map2d.South()}: 1000, // 90deg turn
		{Position: start, Direction: map2d.West()}:  2000, // 180deg turn
	}
	for i, row := range maze {
		for j, tile := range row {
			if tile == Empty || tile == End {
				// 2: all other nodes have infinite (unknown) cost
				pos := map2d.PositionFromIndex(i, j)
				for _, dir := range map2d.CardinalDirections() {
					unvisited[pathNode{Position: pos, Direction: dir}] = math.MaxInt
				}
			}
		}
	}
	return unvisited
}

type pathNode struct {
	// Position of this node in the path, i.e. the tip of the path.
	Position map2d.Position
	// Direction the reindeer is facing as of this node.
	Direction map2d.Direction
}
