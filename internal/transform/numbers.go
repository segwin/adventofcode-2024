package transform

type Float interface {
	float32 | float64
}

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Number interface {
	Float | Signed | Unsigned
}

// Abs returns the absolute value of v. Unlike math.Abs, it accepts any numeric type without conversion.
func Abs[T Number](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
