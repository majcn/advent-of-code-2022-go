package util

import "math"

type Number interface {
	int
}

func Min[T Number](s ...T) (rc T) {
	rc = s[0]

	for _, el := range s[1:] {
		if el < rc {
			rc = el
		}
	}

	return
}

func Max[T Number](s ...T) (rc T) {
	rc = s[0]

	for _, el := range s[1:] {
		if el > rc {
			rc = el
		}
	}

	return
}

func Sum[T Number](s []T) (rc T) {
	for _, el := range s {
		rc += el
	}

	return
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

func Pow[T Number](a T, b T) T {
	return T(math.Pow(float64(a), float64(b)))
}

// Mod function like in Python
func Mod[T Number](a T, b T) T {
	return (a%b + b) % b
}

// SumNaturalNumbers returns sum of natural numbers
// param a: the first term
// param d: the difference between the two consecutive terms
// param n: number of natural numbers
func SumNaturalNumbers(a int, d int, n int) int {
	return n * (2*a + (n-1)*d) / 2
}

// LineIntersection function returns Point where two lines intersects
func LineIntersection(line1 [2]Location, line2 [2]Location) (Location, bool) {
	xdiff := Location{X: line1[0].X - line1[1].X, Y: line2[0].X - line2[1].X}
	ydiff := Location{X: line1[0].Y - line1[1].Y, Y: line2[0].Y - line2[1].Y}

	det := func(a Location, b Location) int {
		return a.X*b.Y - a.Y*b.X
	}

	div := det(xdiff, ydiff)
	if div == 0 {
		return Location{}, false
	}
	d := Location{X: det(line1[0], line1[1]), Y: det(line2[0], line2[1])}
	x := det(d, xdiff) / div
	y := det(d, ydiff) / div

	return Location{X: x, Y: y}, true
}
