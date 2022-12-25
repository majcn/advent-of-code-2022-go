package util

import (
	"os"
	"testing"
)

func InitData[DataType any](parseData func(data string) DataType) DataType {
	dat, _ := os.ReadFile("./input.txt")
	return parseData(string(dat))
}

func AssertTestPartX[R comparable](t *testing.T, expected R, actual R) {
	if actual != expected {
		t.Errorf("Result should be %v, got %v.", expected, actual)
	}
}
