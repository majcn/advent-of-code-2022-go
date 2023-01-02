package util

type BBox struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

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

func (p Point) InBBox(bbox BBox) bool {
	return p.X >= bbox.MinX && p.X <= bbox.MaxX &&
		p.Y >= bbox.MinY && p.Y <= bbox.MaxY
}

type Point3 struct {
	X int
	Y int
	Z int
}

func (p Point3) Add(q Point3) Point3 {
	return Point3{X: p.X + q.X, Y: p.Y + q.Y, Z: p.Z + q.Z}
}

func (p Point3) Mul(k int) Point3 {
	return Point3{X: p.X * k, Y: p.Y * k, Z: p.Z * k}
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
