package main

import (
	"fmt"
	"os"
	"petalv/aoc_2022/file"
	"strconv"
	"strings"
)

type Section struct {
	low  int
	high int
}

type Pair struct {
	id     int
	first  Section
	second Section
}

func main() {
	// part := "part1"
	part := os.Getenv("part")
	input, _ := file.ReadStringArray("input.txt")
	if part == "part2" {
		sum := part2(input)
		fmt.Printf("%d\n", sum)
	} else {
		sum := part1(input)
		fmt.Printf("%d\n", sum)
	}
}

func part1(rawItems []string) int {
	pairs := mapPairs(rawItems)
	sum := 0
	for _, pair := range pairs {
		if contained(pair) {
			// println(pair.id)
			sum++
		}
	}
	return sum
}

func part2(rawItems []string) int {
	pairs := mapPairs(rawItems)
	sum := 0
	for _, pair := range pairs {
		if contained2(pair) {
			// println(pair.id)
			sum++
		}
	}
	return sum
}

func contained(pair Pair) bool {
	var largest Section
	var smallest Section
	if pair.first.high-pair.first.low >= pair.second.high-pair.second.low {
		largest = pair.first
		smallest = pair.second
	} else {
		largest = pair.second
		smallest = pair.first
	}
	return largest.low <= smallest.low && largest.high >= smallest.high
}
func contained2(pair Pair) bool {
	var largest Section
	var smallest Section
	if pair.first.high-pair.first.low >= pair.second.high-pair.second.low {
		largest = pair.first
		smallest = pair.second
	} else {
		largest = pair.second
		smallest = pair.first
	}
	return smallest.low <= largest.high && smallest.high >= largest.low
}
func mapPairs(raw []string) []Pair {
	var pairs []Pair
	for i, row := range raw {
		var pair = Pair{id: i + 1}
		pairSplit := strings.Split(row, ",")
		section1 := strings.Split(pairSplit[0], "-")
		section2 := strings.Split(pairSplit[1], "-")
		s1low, _ := strconv.Atoi(section1[0])
		s1High, _ := strconv.Atoi(section1[1])
		s2low, _ := strconv.Atoi(section2[0])
		s2High, _ := strconv.Atoi(section2[1])
		pair.first = Section{low: s1low, high: s1High}
		pair.second = Section{low: s2low, high: s2High}
		pairs = append(pairs, pair)
	}
	return pairs
}
