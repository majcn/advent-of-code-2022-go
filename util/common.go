package util

type HasLen interface {
	Len() int
}

func SmallerFirst[T HasLen](s1, s2 T) (T, T) {
	if s1.Len() < s2.Len() {
		return s1, s2
	} else {
		return s2, s1
	}
}
