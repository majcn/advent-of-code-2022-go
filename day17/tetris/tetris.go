package tetris

import (
	"hash/fnv"
	"strconv"

	. "majcn.si/advent-of-code-2022/util"
)

type Rock []Point

type Game struct {
	solidRocks         map[int][]int
	movingRock         Rock
	movingRockLocation Point
	maxGridX           int
	maxGridY           int
}

func NewGame(rock Rock) *Game {
	return &Game{
		solidRocks:         make(map[int][]int),
		movingRock:         rock,
		movingRockLocation: Point{X: 2, Y: 3},
		maxGridX:           6,
		maxGridY:           -1,
	}
}

func (g *Game) PlaceRock(rock Rock) {
	g.movingRock = rock
	g.movingRockLocation = Point{X: 2, Y: g.maxGridY + 4}
}

func (g *Game) TransformToSolid() {
	for _, part := range g.movingRock {
		xx := g.movingRockLocation.X + part.X
		yy := g.movingRockLocation.Y + part.Y

		if _, ok := g.solidRocks[yy]; !ok {
			g.solidRocks[yy] = make([]int, g.maxGridX+1)
		}
		g.solidRocks[yy][xx] = 1

		g.maxGridY = Max(g.maxGridY, yy)
	}
}

func (g *Game) Move(direction byte) bool {
	switch direction {
	case '>':
		if g.canMove(g.movingRockLocation.Add(Point{X: 1, Y: 0})) {
			g.movingRockLocation.X++
		}
	case '<':
		if g.canMove(g.movingRockLocation.Add(Point{X: -1, Y: 0})) {
			g.movingRockLocation.X--
		}
	}

	if g.canMove(g.movingRockLocation.Add(Point{X: 0, Y: -1})) {
		g.movingRockLocation.Y--
		return true
	}

	return false
}

func (g *Game) canMove(to Point) bool {
	if to.Y == -1 {
		return false
	}

	for _, part := range g.movingRock {
		xx := part.X + to.X
		if xx == -1 || xx == g.maxGridX+1 {
			return false
		}

		yy := part.Y + to.Y
		if _, ok := g.solidRocks[yy]; ok {
			for x, v := range g.solidRocks[yy] {
				if v == 1 && x == xx {
					return false
				}
			}
		}
	}

	return true
}

func (g *Game) Score() int {
	return g.maxGridY + 1
}

func (g *Game) SolidRocksHashable(size int) uint64 {
	minY := g.maxGridY - size
	if minY <= 0 {
		return 0
	}

	h := fnv.New64()

	for y := minY; y <= g.maxGridY; y++ {
		for x, v := range g.solidRocks[y] {
			if v == 1 {
				h.Write([]byte(strconv.Itoa(x)))
				h.Write([]byte(strconv.Itoa(y - minY)))
			}
		}
	}

	return h.Sum64()
}
