package util

import (
	"math"
	"math/bits"
)

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

func MinAny[T any, R int](s []T, less func(i, j int) bool) (rc T) {
	minIndex := 0

	for i := range s[1:] {
		if less(i+1, minIndex) {
			minIndex = i + 1
		}
	}

	return s[minIndex]
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
func LineIntersection(line1 [2]Point, line2 [2]Point) (Point, bool) {
	xdiff := Point{X: line1[0].X - line1[1].X, Y: line2[0].X - line2[1].X}
	ydiff := Point{X: line1[0].Y - line1[1].Y, Y: line2[0].Y - line2[1].Y}

	det := func(a Point, b Point) int {
		return a.X*b.Y - a.Y*b.X
	}

	div := det(xdiff, ydiff)
	if div == 0 {
		return Point{}, false
	}
	d := Point{X: det(line1[0], line1[1]), Y: det(line2[0], line2[1])}
	x := det(d, xdiff) / div
	y := det(d, ydiff) / div

	return Point{X: x, Y: y}, true
}

// Combinations returns combinations of n elements for a given array.
// For n < 1, it equals to All and returns all combinations.
// source: https://github.com/mxschmitt/golang-combinations/blob/main/combinations.go
func Combinations[T any](set []T, n int) (subsets [][]T) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []T

		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}

		subsets = append(subsets, subset)
	}

	return subsets
}

// Permutations returns permutations of n elements for a given array.
func Permutations[T any](set []T, n int) [][]T {
	if n == 1 {
		temp := make([][]T, 0)
		for _, rr := range set {
			t := make([]T, 0)
			t = append(t, rr)
			temp = append(temp, [][]T{t}...)
		}
		return temp
	}

	res := make([][]T, 0)
	for i := 0; i < len(set); i++ {
		perms := make([]T, 0)
		perms = append(perms, set[:i]...)
		perms = append(perms, set[i+1:]...)
		for _, x := range Permutations(perms, n-1) {
			t := append(x, set[i])
			res = append(res, [][]T{t}...)
		}
	}
	return res
}
