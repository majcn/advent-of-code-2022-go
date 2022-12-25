package util

import (
	"log"
	"strconv"
)

const MaxInt = int(^uint(0) >> 1)

func ParseInt(s any) (rc int) {
	switch sv := s.(type) {
	case byte:
		rc = int(sv - '0')
	case rune:
		rc = int(sv - '0')
	case string:
		rc, _ = strconv.Atoi(sv)
	case []int:
		tmp := make([]byte, len(sv))
		for i, v := range sv {
			tmp[i] = byte('0' + v)
		}
		rc, _ = strconv.Atoi(string(tmp))
	case []byte:
		rc, _ = strconv.Atoi(string(sv))
	default:
		log.Fatalln("Invalid input for ParseInt!")
	}

	return
}

func StringsToInts(s []string) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = ParseInt(v)
	}
	return result
}
