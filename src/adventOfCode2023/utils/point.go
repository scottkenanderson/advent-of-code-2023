package utils

type CoordinateRange struct {
	Start, End Coordinate
}

type Coordinate struct {
	X, Y int
}

func IsInRange(c Coordinate, r CoordinateRange) bool {
	return r.Start.X <= c.X && c.X <= r.End.X &&
		r.Start.Y <= c.Y && c.Y <= r.End.Y
}
