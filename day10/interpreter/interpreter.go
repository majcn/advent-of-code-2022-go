package interpreter

import (
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Interpreter struct {
	program                []string
	programIndex           int
	currentProgramProgress int
	currentCommand         func()
	RegisterX              int
	Cycle                  int
	Halt                   bool
}

func New(program []string) *Interpreter {
	interpreter := &Interpreter{
		program:                program,
		programIndex:           0,
		currentProgramProgress: 0,
		currentCommand:         nil,
		RegisterX:              1,
		Cycle:                  0,
		Halt:                   false,
	}

	interpreter.initNextCommand()
	return interpreter
}

func (interpreter *Interpreter) initNextCommand() {
	command := interpreter.program[interpreter.programIndex]
	commandName, commandValue, _ := strings.Cut(command, " ")

	switch commandName {
	case "noop":
		interpreter.currentProgramProgress = 1
		interpreter.currentCommand = func() {}
	case "addx":
		interpreter.currentProgramProgress = 2
		interpreter.currentCommand = func() {
			interpreter.RegisterX += ParseInt(commandValue)
		}
	}

	interpreter.programIndex++
}

func (interpreter *Interpreter) ExecSingleCycle() {
	if interpreter.Halt {
		return
	}

	interpreter.Cycle++

	if interpreter.currentProgramProgress == 1 {
		interpreter.currentCommand()

		if interpreter.programIndex == len(interpreter.program) {
			interpreter.Halt = true
			return
		}

		interpreter.initNextCommand()
	} else {
		interpreter.currentProgramProgress -= 1
	}
}
