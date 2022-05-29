package square

import "math"

type Point struct {
	X, Y int
}

type Square struct {
	Point
	a uint
}

func (s Square) End() Point {
	length := int(s.a)
	return Point{s.X + length, s.Y + length}
}

func (s Square) Area() uint {
	return uint(math.Sqrt(float64(s.a)))
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
