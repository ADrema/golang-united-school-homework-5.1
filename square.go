package square

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
}

func (s Square) Area() uint {
	length := s.a
	return length * length
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
