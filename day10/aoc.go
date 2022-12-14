package main

import (
	"fmt"
	"os"
	"petalv/aoc_2022/file"
	"strconv"
	"strings"
)

const (
	SYSTEM_UNSET = -1
	ADD          = 0
	NOOP         = 1
)

const (
	REGISTER_X = 2
)

type Command struct {
	ctype    int
	register int
	value    int
}

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStringArray("input.txt")
	if part == "part2" {
		handle(input, true)
		fmt.Printf("%s", "ECZUZALR")
	} else {
		_, checksum := handle(input, false)
		fmt.Printf("%d", checksum)
	}
}

func machine(commands []Command, isCrt bool) (int, int) {
	stack := commands
	clock := 0
	interruptAtCycle := 1
	var registers = make(map[int]int)
	registers[REGISTER_X] = 1
	checksum := 0
	probeAtCycle := 20
	crtCrgAtCycle := 40
	var currentCommand Command = Command{ctype: SYSTEM_UNSET}
	crt := matrix(7, 41)
	crtRow := 0
	crtCol := 0
	for interruptAtCycle != -1 {
		clock++
		if interruptAtCycle == clock {
			interruptAtCycle = -1
			if currentCommand.ctype == ADD {
				registers[currentCommand.register] += currentCommand.value
			}
			if len(stack) > 0 {
				currentCommand, stack = stack[0], stack[1:]
				if currentCommand.ctype == NOOP {
					interruptAtCycle = clock + 1
				} else if currentCommand.ctype == ADD {
					interruptAtCycle = clock + 2
				} else if currentCommand.ctype == SYSTEM_UNSET {
					// noop
				}
			}
		}
		if isCrt {
			if crtCol == registers[REGISTER_X] || crtCol == registers[REGISTER_X]+1 || crtCol == registers[REGISTER_X]-1 {
				crt[crtRow][crtCol] = "#"
			} else {
				crt[crtRow][crtCol] = "."
			}
		}
		crtCol++

		if clock > 0 && clock%crtCrgAtCycle == 0 {
			crtRow++
			crtCol = 0
		}
		if clock > 0 && clock%probeAtCycle == 0 {
			checksum += clock * registers[REGISTER_X]
			probeAtCycle += 40
		}

	}
	if isCrt {
		printMatrix(crt)
	}
	return registers[REGISTER_X], checksum
}

func handle(input []string, isCrt bool) (int, int) {
	var commands []Command
	for _, raw := range input {
		cSplits := strings.Split(raw, " ")
		var command Command
		if strings.HasPrefix(cSplits[0], "noop") {
			command = Command{ctype: NOOP}
		} else if strings.HasPrefix(cSplits[0], "add") {
			var register int
			if cSplits[0][len(cSplits[0])-1] == 'x' {
				register = REGISTER_X
			}
			value, _ := strconv.Atoi(cSplits[1])
			command = Command{ctype: ADD, register: register, value: value}
		}
		commands = append(commands, command)
	}

	return machine(commands, isCrt)
}

func matrix(ymax int, xmax int) [][]string {
	var matrix = make([][]string, ymax, xmax)
	for y := 0; y < ymax; y++ {
		matrix[y] = make([]string, xmax)
	}
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			matrix[y][x] = ""
		}
	}
	return matrix
}

func printMatrix(matrix [][]string) {
	for y := 0; y < len(matrix); y++ {
		fmt.Printf("%d ", y)
		for x := 0; x < len(matrix[0]); x++ {
			fmt.Printf("%s", matrix[y][x])
		}
		fmt.Printf("\n")
	}
}
