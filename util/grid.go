package util

type Location struct {
	X int
	Y int
}

func (p Location) Add(q Location) Location {
	return Location{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Location) Mul(k int) Location {
	return Location{X: p.X * k, Y: p.Y * k}
}

var neighbours8 = []Location{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

var neighbours4 = []Location{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func GetNeighbours4() []Location {
	return neighbours4
}

func GetNeighbours8() []Location {
	return neighbours8
}
