package main

import (
	"encoding/json"
	"fmt"
	"os"
	"petalv/aoc_2022/file"
	"sort"
	"strconv"
)

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStringArray("input.txt")
	if part == "part2" {
		fmt.Printf("%d", -127)
	} else {
		sum := part1(input)
		fmt.Printf("%d", sum)
	}
}

func part1(input []string) int {

	matrix := matrix(input)
	sumTreeLine := (len(matrix[0]) * 2) + ((len(matrix) - 2) * 2)
	visibleTrees := make(map[string]bool)

	for y := 0; y < len(matrix)-1; y++ {
		row := getRow(y, matrix)
		visiblesX := look(row)
		visiblesX = append(visiblesX, lookReverse(row)...)
		for _, x := range visiblesX {
			visibleTrees[strconv.Itoa(y)+strconv.Itoa(x)] = true
		}
	}

	for x := 0; x < len(matrix[0]); x++ {
		col := getCol(x, matrix)
		visiblesY := look(col)
		visiblesY = append(visiblesY, lookReverse(col)...)
		for _, y := range visiblesY {
			visibleTrees[strconv.Itoa(y)+strconv.Itoa(x)] = true
		}
	}

	fmt.Printf("%v\n", Keys(visibleTrees))
	//plot(matrix)
	return len(Keys(visibleTrees)) + sumTreeLine

}

func getCol(col int, matrix [][]int) []int {
	var projection []int
	for y := 0; y < len(matrix); y++ {
		projection = append(projection, matrix[y][col])
	}
	return projection
}

func getRow(row int, matrix [][]int) []int {
	return matrix[row]
}

func look(projection []int) []int {
	var visibles []int
	viewPoint := projection[0]
	for z := 1; z < len(projection)-1; z++ {
		if projection[z] > viewPoint {
			visibles = append(visibles, z)
			viewPoint = projection[z]
		}
	}
	return visibles
}

func lookReverse(projection []int) []int {
	visibles := look(reverse(projection))
	var newVisibles []int
	for _, v := range visibles {
		newVisibles = append(newVisibles, (v*-1)+len(projection)-1)
	}
	fmt.Printf("o: %v n: %v\n", visibles, newVisibles)
	return newVisibles
}

func matrix(lines []string) [][]int {
	xmax := len(lines[0])
	ymax := len(lines)
	var matrix = make([][]int, xmax, ymax)
	for y, line := range lines {
		matrix[y] = make([]int, xmax)
		for x, c := range line {
			i, _ := strconv.Atoi(string(c))
			matrix[y][x] = i
		}
	}
	return matrix
}

func reverse(numbers []int) []int {
	newNumbers := make([]int, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}

func plot(matrix [][]int) {
	for _, row := range matrix {
		s, _ := json.Marshal(row)
		fmt.Printf("%s\n", s)
	}
}

func Keys(myMap map[string]bool) []string {
	keys := make([]string, 0, len(myMap))
	for k, _ := range myMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
