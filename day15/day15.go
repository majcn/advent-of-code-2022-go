package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Range struct {
	Min int
	Max int
}

type Sensor struct {
	Point
	ClosestBeacon Point
}

type DataType []Sensor

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	r := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

	result := make([]Sensor, len(dataSplit))
	for i, line := range dataSplit {
		match := r.FindStringSubmatch(line)
		result[i] = Sensor{
			Point:         Point{X: ParseInt(match[1]), Y: ParseInt(match[2])},
			ClosestBeacon: Point{X: ParseInt(match[3]), Y: ParseInt(match[4])},
		}
	}

	return result
}

func solvePartX(data DataType, y int) []Range {
	ranges := make([]Range, 0, len(data))
	for _, sensor := range data {
		maxDistance := Abs(sensor.X-sensor.ClosestBeacon.X) + Abs(sensor.Y-sensor.ClosestBeacon.Y)
		yDistance := Abs(sensor.Y - y)
		if maxDistance > yDistance {
			diff := maxDistance - yDistance
			ranges = append(ranges, Range{Min: sensor.X - diff, Max: sensor.X + diff})
		}
	}

	if len(ranges) == 0 {
		return []Range{}
	}

	sort.Slice(ranges, func(i, j int) bool { return ranges[i].Min < ranges[j].Min })

	for {
		newRanges := make([]Range, 1, len(ranges))
		newRanges[0] = ranges[0]

		noChanges := true
		for _, r2 := range ranges[1:] {
			r1 := newRanges[len(newRanges)-1]

			newRanges = newRanges[:len(newRanges)-1]

			switch {
			case r2.Min < r1.Min && r1.Max < r2.Max:
				newRanges = append(newRanges, r2)
				noChanges = false
			case r1.Min < r2.Min && r2.Max < r1.Max:
				newRanges = append(newRanges, r1)
				noChanges = false
			case r1.Min == r2.Min && r1.Max == r2.Max:
				newRanges = append(newRanges, r1)
				noChanges = false
			case r1.Max < r2.Min:
				newRanges = append(newRanges, r1, r2)
			case r1.Min == r2.Min:
				newRanges = append(newRanges, Range{Min: r1.Min, Max: Max(r1.Max, r2.Max)})
				noChanges = false
			case r1.Max == r2.Max:
				newRanges = append(newRanges, Range{Min: Min(r1.Min, r2.Min), Max: r1.Max})
				noChanges = false
			case r1.Max == r2.Min:
				newRanges = append(newRanges, Range{Min: r1.Min, Max: r2.Max})
				noChanges = false
			case r2.Min < r1.Max && r1.Min < r2.Min:
				newRanges = append(newRanges, Range{Min: r1.Min, Max: r2.Max})
				noChanges = false
			case r1.Min < r2.Max && r2.Min < r1.Min:
				newRanges = append(newRanges, Range{Min: r2.Min, Max: r1.Max})
				noChanges = false
			}
		}

		if noChanges {
			return newRanges
		}

		ranges = newRanges
	}
}

func solvePart1(data DataType) (rc int) {
	const y = 2000000

	ranges := solvePartX(data, y)
	for _, r := range ranges {
		rc += r.Max - r.Min + 1
	}

	ignoreBeacons := make(Set[Point])
	for _, sensor := range data {
		if sensor.ClosestBeacon.Y == y {
			ignoreBeacons.Add(sensor.ClosestBeacon)
		}
	}
	rc -= len(ignoreBeacons)

	return
}

func solvePart2(data DataType) (rc int) {
	const MinXY, MaxXY = 0, 4000000

	lines := make([][2]Point, 0, len(data)*4)
	for _, sensor := range data {
		maxDistance := Abs(sensor.X-sensor.ClosestBeacon.X) + Abs(sensor.Y-sensor.ClosestBeacon.Y)
		leftPoint := Point{X: sensor.X - maxDistance - 1, Y: sensor.Y}
		rightPoint := Point{X: sensor.X - maxDistance + 1, Y: sensor.Y}

		line1 := [2]Point{leftPoint, leftPoint.Add(Point{X: 1, Y: -1})}
		line2 := [2]Point{leftPoint, leftPoint.Add(Point{X: 1, Y: 1})}
		line3 := [2]Point{rightPoint, rightPoint.Add(Point{X: -1, Y: -1})}
		line4 := [2]Point{rightPoint, rightPoint.Add(Point{X: -1, Y: 1})}

		lines = append(lines, line1, line2, line3, line4)
	}

	interestingPoints := make(Set[int])
	for _, line1 := range lines {
		for _, line2 := range lines {
			if line1 != line2 {
				if point, isValidPoint := LineIntersection(line1, line2); isValidPoint {
					if point.Y >= MinXY && point.Y <= MaxXY {
						interestingPoints.Add(point.Y)
					}
				}
			}
		}
	}

	for y := range interestingPoints {
		ranges := solvePartX(data, y)
		if len(ranges) == 2 {
			if ranges[1].Min-ranges[0].Max == 2 {
				return (ranges[0].Max+1)*MaxXY + y
			}
		}
	}

	return
}

func main() {
	data := parseData(FetchInputData(15))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
