package main

import (
	"fmt"
	"os"
	"petalv/aoc_2022/file"
	"strings"
)

type Rucksack struct {
	c1 string
	c2 string
}

func main() {
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
	rucksacks := mapRucksacks(rawItems)
	prioMap := makePrioMap()
	sum := 0
	for _, r := range rucksacks {
		c := findCommon(r)
		sum += sumChars(c, prioMap)
	}
	return sum
}

func part2(rawItems []string) int {
	rucksacks := mapRucksacks(rawItems)
	prioMap := makePrioMap()
	sum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		batch := rucksacks[i:min(i+3, len(rucksacks))]
		c := findCommons(batch)
		sum += sumChars(strings.Split(c, ""), prioMap)
	}
	return sum
}

func mapRucksacks(rawItems []string) []Rucksack {
	var rucksacks []Rucksack
	for _, items := range rawItems {
		l := len(items)
		c1 := items[:l/2]
		c2 := items[l/2:]
		rucksacks = append(rucksacks, Rucksack{c1: c1, c2: c2})
	}
	return rucksacks
}

func findCommon(rucksack Rucksack) []string {
	set := map[string]struct{}{}
	for _, ci1 := range rucksack.c1 {
		for _, ci2 := range rucksack.c2 {
			if ci1 == ci2 {
				set[string(ci2)] = struct{}{}
			}
		}
	}
	var common []string
	for c := range set {
		common = append(common, c)
	}
	return common
}

func findCommons(rucksack []Rucksack) string {
	//set := map[string]struct{}{}

	var items []string
	for _, r1 := range rucksack {
		items = append(items, r1.c1+r1.c2)
	}
	var commons []string
	commons = append(commons, findCommonsItems(items[0], items[1])...)
	commons = append(commons, findCommonsItems(items[0], items[2])...)
	commons = append(commons, findCommonsItems(items[1], items[2])...)

	combined := strings.Join(commons, "")

	for _, s := range combined {
		if strings.Count(combined, string(s)) == 3 {
			return string(s)
		}
	}
	return "bad"
}

func findCommonsItems(items1 string, items2 string) []string {
	set := map[string]struct{}{}
	for _, i1 := range items1 {
		for _, i2 := range items2 {
			if i1 == i2 {
				set[string(i1)] = struct{}{}
			}
		}
	}
	var common []string
	for c := range set {
		common = append(common, c)
	}
	return common
}
func sumChars(chars []string, prioMap map[string]int) int {
	sum := 0
	for _, s := range chars {
		sum += prioMap[string(s)]
	}
	return sum
}

func makePrioMap() map[string]int {
	var prios = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var prioMap = make(map[string]int)
	for i, s := range prios {
		prioMap[string(s)] = i + 1
	}
	return prioMap
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
