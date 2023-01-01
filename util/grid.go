package util

type Point struct {
	X int
	Y int
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Mul(k int) Point {
	return Point{X: p.X * k, Y: p.Y * k}
}

var neighbors8 = []Point{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

var neighbors4 = []Point{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func GetNeighbors4() []Point {
	return neighbors4
}

func GetNeighbors8() []Point {
	return neighbors8
}
