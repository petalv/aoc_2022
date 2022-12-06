package main

import (
	"fmt"
	"os"
	"petalv/aoc_2022/file"
	"strconv"
	"strings"
)

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStringArray("input.txt")
	if part == "part2" {
		result := crateMover(input, "9001")
		fmt.Printf("%s\n", result)
	} else {
		result := crateMover(input, "9000")
		fmt.Printf("%s\n", result)
	}
}

func crateMover(lines []string, mover string) string {
	sectionSeparatorLineNumber := findStartsWith("\n", lines)

	// read moves
	var moves []string
	for i := sectionSeparatorLineNumber + 1; i < len(lines); i++ {
		line := lines[i]
		line = strings.Replace(line, "move ", "", -1)
		line = strings.Replace(line, "from ", "", -1)
		line = strings.Replace(line, "to ", "", -1)
		line = strings.Replace(line, " ", ";", -1)
		moves = append(moves, line)
	}

	// Read crates
	var stacks = make(map[int][]string)
	var crateEnds []int
	cratesStartLineNumber := sectionSeparatorLineNumber - 2
	for i := cratesStartLineNumber; i >= 0; i-- {
		var crateName string
		stackId := -1
		for j := 0; j < len(lines[i]); j++ {
			s := string(lines[i][j])
			if s == "" {
				continue
			} else if s == "[" {
				crateName = ""
			} else if s == "]" {
				if i == cratesStartLineNumber {
					crateEnds = append(crateEnds, j)
				}
				stackId = findIndexFor(j, crateEnds) + 1
				if stacks[stackId] == nil {
					stacks[stackId] = []string{}
				}
				stacks[stackId] = append([]string{crateName}, stacks[stackId]...)
			} else {
				crateName += s
			}
		}
	}
	// Crate Mover
	for _, move := range moves {
		splits := strings.Split(move, ";")
		mc, _ := strconv.Atoi(string(splits[0]))
		fc, _ := strconv.Atoi(string(splits[1]))
		tc, _ := strconv.Atoi(string(splits[2]))

		if mover == "9000" {
			for i := mc; i > 0; i-- {
				crate := stacks[fc][0]
				stacks[fc] = stacks[fc][1:]
				stacks[tc] = append([]string{crate}, stacks[tc]...)
			}
		}
		if mover == "9001" {
			var crates []string
			crates = append(crates, stacks[fc][0:mc]...)
			stacks[fc] = stacks[fc][mc:]
			crates = append(crates, stacks[tc]...)
			stacks[tc] = crates
		}
	}

	var result []string
	for i := 0; i < len(stacks); i++ {
		result = append(result, stacks[i+1][0])
	}
	return strings.Join(result, "")
}

func findIndexFor(charNum int, ints []int) int {
	for i := len(ints); i > 0; i-- {
		if charNum == ints[i-1] {
			return i - 1
		}
		if charNum > ints[i-1] {
			return i
		}
	}
	return -1
}

func findStartsWith(s string, lines []string) int {
	for i, line := range lines {
		if line == "" {
			return i
		}
	}
	return -1
}
