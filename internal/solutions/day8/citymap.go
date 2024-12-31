package day8

import "github.com/segwin/adventofcode-2024/internal/solutions/map2d"

type Tile byte

const (
	Empty Tile = '.'
)

type CityMap = map2d.Map[Tile]
