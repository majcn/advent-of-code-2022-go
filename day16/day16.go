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
	Time             int
	Location         *Valve
	PressureReleased int
	OpenedValves     Set[*Valve]
}

type ValvePath struct {
	To   *Valve
	Cost int
}

func findAllFinalStates(initState State, nextStatesF func(state State) []State) []State {
	candidates := make([]State, 0)

	var findAllFinalStatesInner func(State)
	findAllFinalStatesInner = func(state State) {
		candidates = append(candidates, state)
		for _, n := range nextStatesF(state) {
			findAllFinalStatesInner(n)
		}
	}

	findAllFinalStatesInner(initState)

	return candidates
}

func getNextStates(state State, valvePaths map[*Valve][]ValvePath, maxTime int) []State {
	result := make([]State, 0, len(valvePaths[state.Location]))

	for _, valvePath := range valvePaths[state.Location] {
		if state.OpenedValves.Contains(valvePath.To) {
			continue
		}

		newTime := state.Time + valvePath.Cost
		if newTime > maxTime {
			continue
		}

		newPressureReleased := state.PressureReleased + (maxTime-newTime)*valvePath.To.FlowRate
		newOpenedValves := make(Set[*Valve], state.OpenedValves.Len()+1)
		for k, v := range state.OpenedValves {
			newOpenedValves[k] = v
		}
		newOpenedValves.Add(valvePath.To)

		result = append(result, State{
			Time:             state.Time + valvePath.Cost,
			Location:         valvePath.To,
			PressureReleased: newPressureReleased,
			OpenedValves:     newOpenedValves,
		})
	}

	return result
}

func solvePartX(data DataType, maxTime int) []State {
	valvePaths := make(map[*Valve][]ValvePath)

	for _, valve1 := range data {
		valvePaths[valve1] = make([]ValvePath, 0, len(data))
		for _, valve2 := range data {
			if valve1 != valve2 && valve2.FlowRate > 0 {
				valvePaths[valve1] = append(valvePaths[valve1], ValvePath{
					To:   valve2,
					Cost: bfs(valve1, valve2) + 1,
				})
			}
		}
	}

	initState := State{
		Time:             0,
		Location:         data["AA"],
		PressureReleased: 0,
		OpenedValves:     make(Set[*Valve]),
	}

	nextStatesF := func(state State) []State { return getNextStates(state, valvePaths, maxTime) }
	return findAllFinalStates(initState, nextStatesF)
}

func solvePart1(data DataType) (rc int) {
	allFinalStates := solvePartX(data, 30)
	maxValueState := MinAny(allFinalStates, func(i, j int) bool {
		return allFinalStates[i].PressureReleased > allFinalStates[j].PressureReleased
	})

	return maxValueState.PressureReleased
}

func solvePart2(data DataType) (rc int) {
	allFinalStates := solvePartX(data, 26)
	sort.Slice(allFinalStates, func(i, j int) bool {
		return allFinalStates[i].PressureReleased > allFinalStates[j].PressureReleased
	})

	for _, s1 := range allFinalStates {
		for _, s2 := range allFinalStates {
			if rc >= s1.PressureReleased+s2.PressureReleased {
				break
			}

			if s1.OpenedValves.Disjoint(s2.OpenedValves) {
				rc = Max(rc, s1.PressureReleased+s2.PressureReleased)
			}
		}
	}

	return
}

func main() {
	data := parseData(FetchInputData(16))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
