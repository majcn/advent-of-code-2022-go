package main

import (
	"fmt"

	"majcn.si/advent-of-code-2022/day17/tetris"
	. "majcn.si/advent-of-code-2022/util"
)

type DataType []byte

func parseData(data string) DataType {
	return DataType(data)
}

type CacheKey struct {
	SolidRocks       uint32
	DirectionIndex   int
	CurrentRockIndex int
}

type CacheValue struct {
	NumberOfSolidRocks int
	Score              int
}

type Cache map[CacheKey]CacheValue

func solvePartX(data DataType, gameDuration int, useCache bool) int {
	rocks := []tetris.Rock{
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
	}

	game := tetris.NewGame(rocks[0])
	numberOfSolidRocks := 0
	currentRockIndex := 1

	cache := make(Cache)
	cacheDisabled := !useCache
	additionalScore := 0

	for i := 0; numberOfSolidRocks != gameDuration; i++ {
		directionIndex := i % len(data)
		didMove := game.Move(data[directionIndex])
		if !didMove {
			game.TransformToSolid()
			game.PlaceRock(rocks[currentRockIndex])
			currentRockIndex = (currentRockIndex + 1) % len(rocks)
			numberOfSolidRocks++

			if cacheDisabled {
				continue
			}

			newCacheKey := CacheKey{
				SolidRocks:       game.SolidRocksHashable(100),
				DirectionIndex:   directionIndex,
				CurrentRockIndex: currentRockIndex,
			}

			newCacheValue := CacheValue{
				NumberOfSolidRocks: numberOfSolidRocks,
				Score:              game.Score(),
			}

			if _, ok := cache[newCacheKey]; ok {
				prevCacheValue := cache[newCacheKey]
				diffNumberOfSolidRocks := newCacheValue.NumberOfSolidRocks - prevCacheValue.NumberOfSolidRocks
				diffScore := newCacheValue.Score - prevCacheValue.Score

				multiplier := (gameDuration - numberOfSolidRocks) / diffNumberOfSolidRocks
				gameDuration -= diffNumberOfSolidRocks * multiplier
				additionalScore = diffScore * multiplier
				cacheDisabled = true
			} else {
				cache[newCacheKey] = newCacheValue
			}
		}
	}

	return game.Score() + additionalScore
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, 2022, false)
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, 1000000000000, true)
}

func main() {
	data := parseData(FetchInputData(17))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
