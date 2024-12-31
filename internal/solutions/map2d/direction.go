package map2d

// Direction for a movement on the map grid.
type Direction struct {
	offset Position
}

func North() Direction { return Direction{offset: Position{Y: -1}} }
func East() Direction  { return Direction{offset: Position{X: 1}} }
func South() Direction { return Direction{offset: Position{Y: 1}} }
func West() Direction  { return Direction{offset: Position{X: -1}} }
