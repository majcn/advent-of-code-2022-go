package main

import (
	"fmt"
	"regexp"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Blueprint struct {
	ID                     int
	OreRobotCostOre        int
	ClayRobotCostOre       int
	ObsidianRobotCostOre   int
	ObsidianRobotCostClay  int
	GeodeRobotCostOre      int
	GeodeRobotCostObsidian int
	MaxOreCost             int
	MaxClayCost            int
	MaxObsidianCost        int
}

type DataType []Blueprint

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	regexString := `^Blueprint (\d+): `
	regexString += `Each ore robot costs (\d+) ore. `
	regexString += `Each clay robot costs (\d+) ore. `
	regexString += `Each obsidian robot costs (\d+) ore and (\d+) clay. `
	regexString += `Each geode robot costs (\d+) ore and (\d+) obsidian.$`
	r := regexp.MustCompile(regexString)

	result := make([]Blueprint, len(dataSplit))
	for i, line := range dataSplit {
		match := r.FindStringSubmatch(line)
		result[i] = Blueprint{
			ID:                     ParseInt(match[1]),
			OreRobotCostOre:        ParseInt(match[2]),
			ClayRobotCostOre:       ParseInt(match[3]),
			ObsidianRobotCostOre:   ParseInt(match[4]),
			ObsidianRobotCostClay:  ParseInt(match[5]),
			GeodeRobotCostOre:      ParseInt(match[6]),
			GeodeRobotCostObsidian: ParseInt(match[7]),
		}
		result[i].MaxOreCost = Max(
			result[i].OreRobotCostOre,
			result[i].ClayRobotCostOre,
			result[i].ObsidianRobotCostOre,
			result[i].GeodeRobotCostOre,
		)
		result[i].MaxClayCost = result[i].ObsidianRobotCostClay
		result[i].MaxObsidianCost = result[i].GeodeRobotCostObsidian
	}

	return result
}

type State struct {
	Time           int
	Ore            int
	Clay           int
	Obsidian       int
	Geode          int
	OreRobots      int
	ClayRobots     int
	ObsidianRobots int
	GeodeRobots    int
}

func getNextStates(blueprint Blueprint, state State) []State {
	if state.Ore >= blueprint.GeodeRobotCostOre && state.Obsidian >= blueprint.GeodeRobotCostObsidian {
		return []State{{
			Time:           state.Time + 1,
			Ore:            state.Ore + state.OreRobots - blueprint.GeodeRobotCostOre,
			Clay:           state.Clay + state.ClayRobots,
			Obsidian:       state.Obsidian + state.ObsidianRobots - blueprint.GeodeRobotCostObsidian,
			Geode:          state.Geode + state.GeodeRobots,
			OreRobots:      state.OreRobots,
			ClayRobots:     state.ClayRobots,
			ObsidianRobots: state.ObsidianRobots,
			GeodeRobots:    state.GeodeRobots + 1,
		}}
	}

	if state.ObsidianRobots < blueprint.MaxObsidianCost && state.Ore >= blueprint.ObsidianRobotCostOre && state.Clay >= blueprint.ObsidianRobotCostClay {
		return []State{{
			Time:           state.Time + 1,
			Ore:            state.Ore + state.OreRobots - blueprint.ObsidianRobotCostOre,
			Clay:           state.Clay + state.ClayRobots - blueprint.ObsidianRobotCostClay,
			Obsidian:       state.Obsidian + state.ObsidianRobots,
			Geode:          state.Geode + state.GeodeRobots,
			OreRobots:      state.OreRobots,
			ClayRobots:     state.ClayRobots,
			ObsidianRobots: state.ObsidianRobots + 1,
			GeodeRobots:    state.GeodeRobots,
		}}
	}

	result := make([]State, 0, 3)

	if state.ClayRobots < blueprint.MaxClayCost && state.Ore >= blueprint.ClayRobotCostOre {
		result = append(result, State{
			Time:           state.Time + 1,
			Ore:            state.Ore + state.OreRobots - blueprint.ClayRobotCostOre,
			Clay:           state.Clay + state.ClayRobots,
			Obsidian:       state.Obsidian + state.ObsidianRobots,
			Geode:          state.Geode + state.GeodeRobots,
			OreRobots:      state.OreRobots,
			ClayRobots:     state.ClayRobots + 1,
			ObsidianRobots: state.ObsidianRobots,
			GeodeRobots:    state.GeodeRobots,
		})
	}

	if state.OreRobots < blueprint.MaxOreCost && state.Ore >= blueprint.OreRobotCostOre {
		result = append(result, State{
			Time:           state.Time + 1,
			Ore:            state.Ore + state.OreRobots - blueprint.OreRobotCostOre,
			Clay:           state.Clay + state.ClayRobots,
			Obsidian:       state.Obsidian + state.ObsidianRobots,
			Geode:          state.Geode + state.GeodeRobots,
			OreRobots:      state.OreRobots + 1,
			ClayRobots:     state.ClayRobots,
			ObsidianRobots: state.ObsidianRobots,
			GeodeRobots:    state.GeodeRobots,
		})
	}

	result = append(result, State{
		Time:           state.Time + 1,
		Ore:            state.Ore + state.OreRobots,
		Clay:           state.Clay + state.ClayRobots,
		Obsidian:       state.Obsidian + state.ObsidianRobots,
		Geode:          state.Geode + state.GeodeRobots,
		OreRobots:      state.OreRobots,
		ClayRobots:     state.ClayRobots,
		ObsidianRobots: state.ObsidianRobots,
		GeodeRobots:    state.GeodeRobots,
	})

	return result
}

func canBeDiscarded(state State, bestResult State, blueprint Blueprint, maxTime int) bool {
	diffTime := maxTime - state.Time

	calculatedStateClay := state.Clay +
		SumNaturalNumbers(state.ClayRobots, 1, diffTime)

	maxAdditionalObsidianRobots := Min(calculatedStateClay/blueprint.ObsidianRobotCostClay, diffTime)
	calculatedStateObsidian := state.Obsidian +
		SumNaturalNumbers(state.ObsidianRobots, 1, maxAdditionalObsidianRobots) +
		(diffTime-maxAdditionalObsidianRobots)*(state.ObsidianRobots+maxAdditionalObsidianRobots)

	maxAdditionalGeodeRobots := Min(calculatedStateObsidian/blueprint.GeodeRobotCostObsidian, diffTime)
	calculatedStateGeode := state.Geode +
		SumNaturalNumbers(state.GeodeRobots, 1, maxAdditionalGeodeRobots) +
		(diffTime-maxAdditionalGeodeRobots)*(state.GeodeRobots+maxAdditionalGeodeRobots)

	return calculatedStateGeode <= bestResult.Geode
}

func dfs(initState State, blueprint Blueprint, maxTime int, nextStatesF func(blueprint Blueprint, state State) []State) State {
	bestResult := State{}
	discovered := make(Set[State])

	var dfsInner func(State)
	dfsInner = func(state State) {
		discovered.Add(state)

		for _, n := range nextStatesF(blueprint, state) {
			if discovered.Contains(n) {
				continue
			}

			if n.Time == maxTime {
				if bestResult.Geode < n.Geode {
					bestResult = n
				}
				continue
			}

			if canBeDiscarded(n, bestResult, blueprint, maxTime) {
				continue
			}

			dfsInner(n)
		}
	}

	dfsInner(initState)

	return bestResult
}

func findMaxGeodes(blueprint Blueprint, maxTime int) int {
	return dfs(State{OreRobots: 1}, blueprint, maxTime, getNextStates).Geode
}

func solvePart1(data DataType) (rc int) {
	findMaxGeodes(data[3], 24)
	for _, blueprint := range data {
		rc += blueprint.ID * findMaxGeodes(blueprint, 24)
	}
	return
}

func solvePart2(data DataType) (rc int) {
	rc = 1
	for _, blueprint := range data[:3] {
		rc *= findMaxGeodes(blueprint, 32)
	}
	return
}

func main() {
	data := parseData(FetchInputData(19))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
