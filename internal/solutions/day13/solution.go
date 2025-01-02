package day13

import (
	"fmt"
	"math"

	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Solution struct {
	Machines []ClawMachine
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 13:\n")

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Total cost: %d\n", WinAllPossible(s.Machines))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Total cost: %d\n", WinAllPossible(CorrectMachines(1e13, s.Machines...)))

	return nil
}

// CorrectMachines returns a new set of claw machines with prize positions adjusted by the given offset.
func CorrectMachines(prizeOffset int, machines ...ClawMachine) []ClawMachine {
	correctedMachines := make([]ClawMachine, len(machines))
	for i := range correctedMachines {
		correctedMachines[i] = ClawMachine{
			MoveA: machines[i].MoveA,
			MoveB: machines[i].MoveB,
			Prize: machines[i].Prize.Add(prizeOffset, prizeOffset),
		}
	}
	return correctedMachines
}

// WinAllPossible returns the cost to win all of the given claw machines.
func WinAllPossible(machines []ClawMachine) int {
	totalCost := 0
	for _, m := range machines {
		a, b, ok := OptimalPresses(m)
		if !ok {
			continue // unwinnable
		}
		totalCost += Cost(a, b)
	}
	return totalCost
}

// Cost returns the number of tokens required to press A and B the given number of times.
func Cost(a, b int) int {
	return a*ACost + b*BCost
}

// OptimalPresses returns the lowest-cost number of button presses required to reach the prize.
// If no such sequence is found, ok is set to false.
func OptimalPresses(m ClawMachine) (a, b int, ok bool) {
	// case 1: Px / Py == Ax / Ay and/or Px/Py == Bx / By
	//   -> in this case, A and B both individually allow reaching P and we just need to optimize for cost
	scaleA, canScaleA := m.MoveA.ScalingFactor(map2d.Distance(m.Prize))
	scaleB, canScaleB := m.MoveB.ScalingFactor(map2d.Distance(m.Prize))

	switch {
	case canScaleA && canScaleB:
		// either move works, select cheapest between them
		normalizedACost := costPerMovement(m.MoveA, ACost)
		normalizedBCost := costPerMovement(m.MoveB, BCost)
		if normalizedACost < normalizedBCost {
			return scaleA, 0, true // A is cheaper overall
		}
		return 0, scaleB, true // B is cheaper overall

	case canScaleA:
		return scaleA, 0, true // use A

	case canScaleB:
		return 0, scaleB, true // use B
	}

	// case 2: 2-variable & 2-equation system => 1 unique solution
	//   -> invalid solutions: non-integer press counts
	//
	// equations:
	//   (1) Ax*a + By*b = Px => a = (Px - Bx*b) / Ax
	//   (2) Ay*a + By*b = Py => a = (Py - By*b) / Ay
	//
	// => (Px - Bx*b) / Ax = (Py - By*b) / Ay
	// => Ay*Px - Ay*Bx*b = Ax*Py - Ax*By*b
	// => (3) b = (Ax*Py - Ay*Px) / (Ax*By - Ay*Bx)

	maybeB := float64(m.MoveA.X*m.Prize.Y-m.MoveA.Y*m.Prize.X) / float64(m.MoveA.X*m.MoveB.Y-m.MoveA.Y*m.MoveB.X) // eqn (3)
	if !isWhole(maybeB) || maybeB < 0 {
		return 0, 0, false // no sequence allows us to win this game
	}

	b = int(maybeB)
	maybeA := float64(m.Prize.X-m.MoveB.X*b) / float64(m.MoveA.X) // eqn (1)
	if !isWhole(maybeA) || maybeA < 0 {
		return 0, 0, false // no sequence allows us to win this game
	}

	return int(maybeA), b, true
}

func isWhole(v float64) bool {
	return v == math.Trunc(v)
}

// costPerMovement returns the cost normalized to the amount of movement being produced.
func costPerMovement(amount map2d.Distance, triggerCost int) float64 {
	return float64(triggerCost) / amount.Norm()
}
