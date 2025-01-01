package map2d

import "math"

// Distance on a 2D map.
type Distance Position

// Norm returns the Euclidean norm of this distance.
func (d Distance) Norm() float64 {
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

// ScalingFactor returns the integer scaling factor required to turn d into d2.
// If no such factor exists (d2 is not an integer multiple of d), ok is set to false.
func (d Distance) ScalingFactor(d2 Distance) (scale int, ok bool) {
	if d.X%d2.X != 0 || d.Y%d2.Y != 0 {
		return 0, false // one or both axes is not an integer multiple of d2
	}

	scaleX := d.X / d2.X
	if scaleY := d.Y / d2.Y; scaleX != scaleY {
		return 0, false // not a multiple of d2
	}

	return scaleX, true
}

// Scale d by the given factor as a new Distance object.
func (d Distance) Scale(by int) Distance {
	return Distance{X: d.X * by, Y: d.Y * by}
}
