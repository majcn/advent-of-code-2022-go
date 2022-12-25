package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Valve struct {
	Name     string
	FlowRate int
	Tunnels  []*Valve
}

type DataType map[string]*Valve

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	r := regexp.MustCompile(`Valve ([A-Z][A-Z]) has flow rate=(\d+); tunnels? leads? to valves? ((?:[A-Z][A-Z](?:, )?)+)$`)

	result := make(map[string]*Valve, len(dataSplit))
	for _, line := range dataSplit {
		match := r.FindStringSubmatch(line)

		result[match[1]] = &Valve{}
	}

	for _, line := range dataSplit {
		match := r.FindStringSubmatch(line)

		tunnelNames := strings.Split(match[3], ", ")
		tunnels := make([]*Valve, len(tunnelNames))
		for i, n := range tunnelNames {
			tunnels[i] = result[n]
		}

		valve := result[match[1]]
		valve.Name = match[1]
		valve.FlowRate = ParseInt(match[2])
		valve.Tunnels = tunnels
	}

	return result
}

func bfs(startNode *Valve, endNode *Valve) int {
	type QueueNode struct {
		Node *Valve
		Time int
	}

	queue := make([]QueueNode, 0)
	visited := make(Set[*Valve])

	queue = append(queue, QueueNode{Node: startNode, Time: 0})
	visited.Add(startNode)
	for len(queue) > 0 {
		queueEl := queue[0]
		queue = queue[1:]

		node := queueEl.Node

		if node == endNode {
			return queueEl.Time
		}

		for _, neighbor := range node.Tunnels {
			if !visited.Contains(neighbor) {
				queue = append(queue, QueueNode{Node: neighbor, Time: queueEl.Time + 1})
				visited.Add(neighbor)
			}
		}
	}

	return MaxInt
}

type State struct {
	Time1            int
	Time2            int
	Location1        *Valve
	Location2        *Valve
	PressureReleased int
	OpenedValves     Set[*Valve]
}

type HashableState struct {
	Time1           int
	Time2           int
	Location1       *Valve
	Location2       *Valve
	OpenedValvesStr string
}

func (state *State) Hashable() HashableState {
	tmp := make([]string, len(state.OpenedValves))
	i := 0
	for v := range state.OpenedValves {
		tmp[i] = v.Name
		i++
	}

	sort.Strings(tmp)

	return HashableState{
		Time1:           state.Time1,
		Time2:           state.Time2,
		Location1:       state.Location1,
		Location2:       state.Location2,
		OpenedValvesStr: strings.Join(tmp, ""),
	}
}

func dfs(initState State, nextStatesF func(state State) []State, canBeDiscardedF func(n State, bestResult int) bool) int {
	bestResult := 0
	discovered := make(map[HashableState]int)

	var dfsInner func(State)
	dfsInner = func(state State) {
		discovered[state.Hashable()] = state.PressureReleased

		for _, n := range nextStatesF(state) {
			if value, ok := discovered[n.Hashable()]; ok {
				if state.PressureReleased < value {
					continue
				}
			}

			if bestResult < n.PressureReleased {
				println(bestResult)
				bestResult = n.PressureReleased
			}

			if canBeDiscardedF(n, bestResult) {
				continue
			}

			dfsInner(n)
		}
	}

	dfsInner(initState)

	return bestResult
}

type ValvePath struct {
	To   *Valve
	Cost int
}

func canBeDiscarded(state State, bestResult int, maxTime int, valvePaths ValvePaths) bool {
	bestPossibleValue := 0

	flows := make([]int, 0, len(valvePaths.Paths))
	for valve := range valvePaths.Paths {
		if !state.OpenedValves.Contains(valve) {
			flows = append(flows, valve.FlowRate)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(flows)))
	i := Min(state.Time1, state.Time2)
	for _, flow := range flows {
		bestPossibleValue += flow * (maxTime - i)
		i++

		if maxTime-i == 0 {
			break
		}
	}
	return state.PressureReleased+bestPossibleValue <= bestResult
}

type ValvePaths struct {
	Paths   map[*Valve][]ValvePath
	MaxCost map[*Valve]int
}

func solvePartX(data DataType, maxTime int, nextStatesF func(state State, valvePaths ValvePaths) []State) int {
	valvePaths := ValvePaths{Paths: make(map[*Valve][]ValvePath), MaxCost: make(map[*Valve]int)}

	for _, valve1 := range data {
		valvePaths.Paths[valve1] = make([]ValvePath, 0, len(data))
		for _, valve2 := range data {
			if valve1 != valve2 && valve2.FlowRate > 0 {
				valvePaths.Paths[valve1] = append(valvePaths.Paths[valve1], ValvePath{
					To:   valve2,
					Cost: bfs(valve1, valve2) + 1,
				})
			}
		}
	}

	for valve, paths := range valvePaths.Paths {
		sort.Slice(paths, func(i, j int) bool {
			iv := (maxTime - paths[i].Cost) * paths[i].To.FlowRate
			jv := (maxTime - paths[j].Cost) * paths[j].To.FlowRate
			return iv > jv
		})
		valvePaths.MaxCost[valve] = paths[0].Cost
	}

	initState := State{
		Time1:            0,
		Time2:            0,
		Location1:        data["AA"],
		Location2:        data["AA"],
		PressureReleased: 0,
		OpenedValves:     make(Set[*Valve]),
	}

	return dfs(initState,
		func(state State) []State {
			return nextStatesF(state, valvePaths)
		},
		func(state State, bestResult int) bool {
			return canBeDiscarded(state, bestResult, maxTime, valvePaths)
		},
	)
}

func solvePart1(data DataType) (rc int) {
	maxTime := 30
	nextStatesF := func(state State, valvePaths ValvePaths) []State {
		result := make([]State, 0, len(valvePaths.Paths[state.Location1]))

		for _, valvePath := range valvePaths.Paths[state.Location1] {
			if state.OpenedValves.Contains(valvePath.To) {
				continue
			}

			newTime := state.Time1 + valvePath.Cost
			if newTime > maxTime {
				continue
			}

			newPressureReleased := state.PressureReleased + (maxTime-newTime)*valvePath.To.FlowRate
			newOpenedValves := make(Set[*Valve], len(state.OpenedValves)+1)
			for k, v := range state.OpenedValves {
				newOpenedValves[k] = v
			}
			newOpenedValves.Add(valvePath.To)

			result = append(result, State{
				Time1:            state.Time1 + valvePath.Cost,
				Time2:            maxTime,
				Location1:        valvePath.To,
				Location2:        state.Location2,
				PressureReleased: newPressureReleased,
				OpenedValves:     newOpenedValves,
			})
		}

		return result
	}

	return solvePartX(data, maxTime, nextStatesF)
}

func solvePart2(data DataType) (rc int) {
	maxTime := 26
	nextStatesF := func(state State, valvePaths ValvePaths) []State {
		result := make([]State, 0, len(valvePaths.Paths[state.Location1])*len(valvePaths.Paths[state.Location2]))

		for _, valvePath1 := range valvePaths.Paths[state.Location1] {
			if state.OpenedValves.Contains(valvePath1.To) {
				continue
			}

			for _, valvePath2 := range valvePaths.Paths[state.Location2] {
				if valvePath1.To == valvePath2.To {
					continue
				}

				if state.OpenedValves.Contains(valvePath2.To) {
					continue
				}

				newTime1 := state.Time1 + valvePath1.Cost
				newTime2 := state.Time2 + valvePath2.Cost
				if newTime1 > maxTime || newTime2 > maxTime {
					continue
				}

				newPressureReleased := state.PressureReleased + (maxTime-newTime1)*valvePath1.To.FlowRate + (maxTime-newTime2)*valvePath2.To.FlowRate
				newOpenedValves := make(Set[*Valve], len(state.OpenedValves)+2)
				for k, v := range state.OpenedValves {
					newOpenedValves[k] = v
				}
				newOpenedValves.Add(valvePath1.To)
				newOpenedValves.Add(valvePath2.To)

				result = append(result, State{
					Time1:            newTime1,
					Time2:            newTime2,
					Location1:        valvePath1.To,
					Location2:        valvePath2.To,
					PressureReleased: newPressureReleased,
					OpenedValves:     newOpenedValves,
				})
			}
		}

		if state.Time2+valvePaths.MaxCost[state.Location2] > maxTime {
			for _, valvePath1 := range valvePaths.Paths[state.Location1] {
				if state.OpenedValves.Contains(valvePath1.To) {
					continue
				}

				newTime1 := state.Time1 + valvePath1.Cost
				if newTime1 > maxTime {
					continue
				}

				newPressureReleased := state.PressureReleased + (maxTime-newTime1)*valvePath1.To.FlowRate
				newOpenedValves := make(Set[*Valve], len(state.OpenedValves)+1)
				for k, v := range state.OpenedValves {
					newOpenedValves[k] = v
				}
				newOpenedValves.Add(valvePath1.To)

				result = append(result, State{
					Time1:            newTime1,
					Time2:            state.Time2,
					Location1:        valvePath1.To,
					Location2:        state.Location2,
					PressureReleased: newPressureReleased,
					OpenedValves:     newOpenedValves,
				})
			}
		}

		if state.Time1+valvePaths.MaxCost[state.Location1] > maxTime {
			for _, valvePath2 := range valvePaths.Paths[state.Location2] {
				if state.OpenedValves.Contains(valvePath2.To) {
					continue
				}

				newTime2 := state.Time2 + valvePath2.Cost
				if newTime2 > maxTime {
					continue
				}

				newPressureReleased := state.PressureReleased + (maxTime-newTime2)*valvePath2.To.FlowRate
				newOpenedValves := make(Set[*Valve], len(state.OpenedValves)+1)
				for k, v := range state.OpenedValves {
					newOpenedValves[k] = v
				}
				newOpenedValves.Add(valvePath2.To)

				result = append(result, State{
					Time1:            state.Time1,
					Time2:            newTime2,
					Location1:        state.Location1,
					Location2:        valvePath2.To,
					PressureReleased: newPressureReleased,
					OpenedValves:     newOpenedValves,
				})
			}
		}

		return result
	}

	return solvePartX(data, maxTime, nextStatesF)
}

func main() {
	data := parseData(FetchInputData(16))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
