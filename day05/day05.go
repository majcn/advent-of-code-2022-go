package main

import (
	"fmt"
	"regexp"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Command struct {
	N    int
	From int
	To   int
}

type Stack []byte
type State []Stack

type DataType struct {
	State    State
	Commands []Command
}

func parseData(data string) DataType {
	stateLines, commandLines, _ := strings.Cut(data, "\n\n")

	commandsRegex := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	commandsDataSplit := strings.Split(commandLines, "\n")
	commands := make([]Command, len(commandsDataSplit))
	for i, line := range commandsDataSplit {
		match := commandsRegex.FindStringSubmatch(line)
		commands[i] = Command{
			N:    ParseInt(match[1]),
			From: ParseInt(match[2]) - 1,
			To:   ParseInt(match[3]) - 1,
		}
	}

	stateDataSplit := strings.Split(stateLines, "\n")
	lastLine := stateDataSplit[len(stateDataSplit)-1]
	state := make([]Stack, (len(lastLine)+1)/4)
	for i := 0; i < len(lastLine); i += 4 {
		state[i/4] = make(Stack, 0)
	}

	for lineNumber := len(stateDataSplit) - 2; lineNumber >= 0; lineNumber-- {
		line := stateDataSplit[lineNumber]
		for i, j := 0, 0; i < len(line); i, j = i+4, j+1 {
			if line[i+1] != ' ' {
				state[j] = append(state[j], line[i+1])
			}
		}
	}

	return DataType{
		State:    state,
		Commands: commands,
	}
}

func (state State) DeepCopy() State {
	result := make(State, len(state))
	for i := range state {
		result[i] = make(Stack, len(state[i]))
		copy(result[i], state[i])
	}
	return result
}

func (stack *Stack) Pop(n int) Stack {
	tmp := *stack
	result := tmp[len(tmp)-n:]
	*stack = tmp[:len(*stack)-n]

	return result
}

func solvePartX(data DataType, applyCommand func(state []Stack, command Command)) string {
	state, commands := data.State.DeepCopy(), data.Commands

	for _, command := range commands {
		applyCommand(state, command)
	}

	result := make([]byte, len(state))
	for i, crates := range state {
		result[i] = crates[len(crates)-1]
	}
	return string(result)
}

func solvePart1(data DataType) (rc string) {
	applyCommand := func(state []Stack, command Command) {
		crates := state[command.From].Pop(command.N)
		for i := len(crates) - 1; i >= 0; i-- {
			state[command.To] = append(state[command.To], crates[i])
		}
	}

	return solvePartX(data, applyCommand)
}

func solvePart2(data DataType) (rc string) {
	applyCommand := func(state []Stack, command Command) {
		crates := state[command.From].Pop(command.N)
		state[command.To] = append(state[command.To], crates...)
	}

	return solvePartX(data, applyCommand)
}

func main() {
	data := parseData(FetchInputData(5))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
