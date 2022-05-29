package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	length := int(s.a)
	return Point{s.start.x + length, s.start.y + length}
	// implement me
}

func (s Square) Area() uint {
	return uint(math.Sqrt(float64(s.a)))
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
