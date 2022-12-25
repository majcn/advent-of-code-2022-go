package util

func StringPermutations(L []string, r int) [][]string {
	if r == 1 {
		temp := make([][]string, 0)
		for _, rr := range L {
			t := make([]string, 0)
			t = append(t, rr)
			temp = append(temp, [][]string{t}...)
		}
		return temp
	}

	res := make([][]string, 0)
	for i := 0; i < len(L); i++ {
		perms := make([]string, 0)
		perms = append(perms, L[:i]...)
		perms = append(perms, L[i+1:]...)
		for _, x := range StringPermutations(perms, r-1) {
			t := append(x, L[i])
			res = append(res, [][]string{t}...)
		}
	}
	return res
}

func MapKeysAsStringSet(m map[string]string) Set[string] {
	result := make(Set[string], len(m))
	for k := range m {
		result[k] = voidVar
	}

	return result
}

func MapValuesAsStringSet(m map[string]string) Set[string] {
	result := make(Set[string], len(m))
	for _, v := range m {
		result[v] = voidVar
	}

	return result
}

func Disjoint(as1 []string, as2 []string) bool {
	for _, s1 := range as1 {
		for _, s2 := range as2 {
			if s1 == s2 {
				return false
			}
		}
	}

	return true
}
