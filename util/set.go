package util

type Set[K comparable] map[K]void

func NewSet[K comparable](s []K) Set[K] {
	result := make(Set[K], len(s))
	for _, v := range s {
		result[v] = voidVar
	}

	return result
}

func (s *Set[K]) Add(el K) {
	(*s)[el] = voidVar
}

func (s *Set[K]) Remove(el K) {
	delete(*s, el)
}

func (s *Set[K]) Pop() K {
	for v := range *s {
		s.Remove(v)
		return v
	}

	var result K
	return result
}

func (s *Set[K]) Contains(el K) bool {
	if _, ok := (*s)[el]; ok {
		return true
	}

	return false
}

func (s *Set[K]) IsSubset(s2 *Set[K]) bool {
	for el := range *s {
		if !s2.Contains(el) {
			return false
		}
	}

	return true
}

func (s *Set[K]) Equals(s2 *Set[K]) bool {
	if len(*s) != len(*s2) {
		return false
	}

	return s.IsSubset(s2)
}

func (s *Set[K]) Intersection(s2 *Set[K]) Set[K] {
	result := make(Set[K], len(*s))

	for el := range *s {
		if (*s2).Contains(el) {
			result[el] = voidVar
		}
	}

	return result
}

func (s *Set[K]) Difference(s2 *Set[K]) Set[K] {
	result := make(Set[K], len(*s))

	for el := range *s {
		if !(*s2).Contains(el) {
			result[el] = voidVar
		}
	}

	return result
}
