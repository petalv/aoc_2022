package main

import (
	"fmt"
	"math"
	"os"
	"petalv/aoc_2022/file"
	"strconv"
	"strings"
)

const (
	RIGHT = 0
	LEFT  = 1
	UP    = 2
	DOWN  = 3
)

type Move struct {
	direction int
	steps     int
}

type Pos struct {
	y int
	x int
}
type Head struct {
	pos     *Pos
	tail    *Tail
	visited []Pos
}

type Tail struct {
	pos     *Pos
	tail    *Tail
	visited []Pos
}

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStringArray("input.txt")
	if part == "part2" {
		sum := part2(input)
		fmt.Printf("%d", sum)
	} else {
		sum := part1(input)
		fmt.Printf("%d", sum)
	}
}

func part2(input []string) int {

	var moves []Move
	for _, s := range input {
		splits := strings.Split(s, " ")
		steps, _ := strconv.Atoi(splits[1])
		direction := getDirection(splits[0])
		moves = append(moves, Move{direction: direction, steps: steps})
	}

	// matrix := matrix(100, 200)
	tail9 := &Tail{pos: &Pos{y: 1000, x: 1000}}
	tail8 := &Tail{pos: &Pos{y: 1000, x: 1000}, tail: tail9}
	tail7 := &Tail{pos: &Pos{y: 1000, x: 1000}, tail: tail8}
	tail6 := &Tail{pos: &Pos{y: 1000, x: 1000}, tail: tail7}
	tail5 := &Tail{pos: &Pos{y: 1000, x: 1000}, tail: tail6}
	tail4 := &Tail{pos: &Pos{y: 1000, x: 1000}, tail: tail5}
	tail3 := &Tail{pos: &Pos{y: 1000, x: 1000}, tail: tail4}
	tail2 := &Tail{pos: &Pos{y: 1000, x: 1000}, tail: tail3}
	tail1 := &Tail{pos: &Pos{y: 1000, x: 1000}, tail: tail2}
	head := &Head{pos: &Pos{y: 1000, x: 1000}, tail: tail1}

	head.visited = append(head.visited, *head.pos)
	tail1.visited = append(head.tail.visited, *tail1.pos)
	tail2.visited = append(head.tail.visited, *tail2.pos)
	tail3.visited = append(head.tail.visited, *tail3.pos)
	tail4.visited = append(head.tail.visited, *tail4.pos)
	tail5.visited = append(head.tail.visited, *tail5.pos)
	tail6.visited = append(head.tail.visited, *tail6.pos)
	tail7.visited = append(head.tail.visited, *tail7.pos)
	tail8.visited = append(head.tail.visited, *tail8.pos)
	tail9.visited = append(head.tail.visited, *tail9.pos)

	for _, m := range moves {
		head = step(head, m)
		// plot(matrix, head.visited, "H")
		// plot(matrix, tail9.visited, "#")
		// printMatrix(matrix)
	}
	//return len(zapCoords(head.tail.visited))
	return len(zapCoords(tail9.visited))
}

func part1(input []string) int {

	var moves []Move
	for _, s := range input {
		splits := strings.Split(s, " ")
		steps, _ := strconv.Atoi(splits[1])
		direction := getDirection(splits[0])
		moves = append(moves, Move{direction: direction, steps: steps})
	}

	// matrix := matrix(100, 200)
	tail1 := &Tail{pos: &Pos{y: 1000, x: 1000}}
	head := &Head{pos: &Pos{y: 1000, x: 1000}, tail: tail1}

	head.visited = append(head.visited, *head.pos)
	tail1.visited = append(head.tail.visited, *tail1.pos)

	for _, m := range moves {
		head = step(head, m)
		// plot(matrix, head.visited, "H")
		// plot(matrix, tail9.visited, "#")
		// printMatrix(matrix)
	}
	return len(zapCoords(head.tail.visited))
}

func plot(matrix [][]string, plots []Pos, marker string) {
	for _, p := range plots {
		matrix[p.y][p.x] = marker
	}
}

func getDirection(dirStr string) int {
	switch dirStr {
	case "R":
		return RIGHT
	case "L":
		return LEFT
	case "D":
		return DOWN
	case "U":
		return UP
	}
	return -1
}

func step(head *Head, move Move) *Head {
	for i := 0; i < move.steps; i++ {
		// fmt.Printf("[%d,%d] [%d,%d]\n", head.pos.y, head.pos.x, head.tail.pos.y, head.tail.pos.x)
		if move.direction == RIGHT {
			head.pos.x++
			if euclidianDistance(*head.pos, *head.tail.pos) == 2 {
				if head.pos.y == head.tail.pos.y {
					head.tail.pos.x++
				} else {
					if head.pos.y > head.tail.pos.y {
						head.tail.pos.y++
						head.tail.pos.x++
					} else {
						head.tail.pos.y--
						head.tail.pos.x++
					}
				}
			}
		} else if move.direction == LEFT {
			head.pos.x--
			if euclidianDistance(*head.pos, *head.tail.pos) == 2 {
				if head.pos.y == head.tail.pos.y {
					head.tail.pos.x--
				} else {
					if head.pos.y > head.tail.pos.y {
						head.tail.pos.y++
						head.tail.pos.x--
					} else {
						head.tail.pos.y--
						head.tail.pos.x--
					}
				}
			}
		} else if move.direction == UP {
			head.pos.y--
			if euclidianDistance(*head.pos, *head.tail.pos) == 2 {
				if head.pos.x == head.tail.pos.x {
					head.tail.pos.y--
				} else {
					if head.pos.x > head.tail.pos.x {
						head.tail.pos.x++
						head.tail.pos.y--
					} else {
						head.tail.pos.x--
						head.tail.pos.y--
					}
				}
			}
		} else if move.direction == DOWN {
			head.pos.y++
			if euclidianDistance(*head.pos, *head.tail.pos) == 2 {
				if head.pos.x == head.tail.pos.x {
					head.tail.pos.y++
				} else {
					if head.pos.x > head.tail.pos.x {
						head.tail.pos.x++
						head.tail.pos.y++
					} else {
						head.tail.pos.x--
						head.tail.pos.y++
					}
				}
			}
		}
		head.tail.visited = append(head.tail.visited, *head.tail.pos)
		head.visited = append(head.visited, *head.pos)
		rattleTail(head.tail)
	}
	// fmt.Printf("[%d,%d] [%d,%d]\n", head.pos.y, head.pos.x, head.tail.pos.y, head.tail.pos.x)
	return head
}

func rattleTail(tail *Tail) {
	if tail.tail != nil {
		if euclidianDistance(*tail.pos, *tail.tail.pos) == 2 {
			if tail.pos.y > tail.tail.pos.y {
				tail.tail.pos.y++
				if tail.pos.x > tail.tail.pos.x {
					tail.tail.pos.x++
				} else if tail.pos.x < tail.tail.pos.x {
					tail.tail.pos.x--
				}
			} else if tail.pos.y < tail.tail.pos.y {
				tail.tail.pos.y--
				if tail.pos.x > tail.tail.pos.x {
					tail.tail.pos.x++
				} else if tail.pos.x < tail.tail.pos.x {
					tail.tail.pos.x--
				}
			} else if tail.pos.x < tail.tail.pos.x {
				tail.tail.pos.x--
				if tail.pos.y > tail.tail.pos.y {
					tail.tail.pos.y++
				} else if tail.pos.y < tail.tail.pos.y {
					tail.tail.pos.y--
				}
			} else if tail.pos.x > tail.tail.pos.x {
				tail.tail.pos.x++
				if tail.pos.y > tail.tail.pos.y {
					tail.tail.pos.y++
				} else if tail.pos.y < tail.tail.pos.y {
					tail.tail.pos.y--
				}
			}

		}
		tail.tail.visited = append(tail.tail.visited, *tail.tail.pos)
		rattleTail(tail.tail)
	}
}

func manhattanDistance(pos Pos, pos2 Pos) int {
	return absInt(pos.x-pos2.x) + absInt(pos.y-pos2.y)
}
func euclidianDistance(pos Pos, pos2 Pos) int {
	var xd = absInt(pos.x - pos2.x)
	var yd = absInt(pos.y - pos2.y)
	return int(math.Floor(math.Sqrt(float64(xd*xd + yd*yd))))
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

func matrix(ymax int, xmax int) [][]string {
	var matrix = make([][]string, ymax, xmax)
	for y := 0; y < ymax; y++ {
		matrix[y] = make([]string, xmax)
	}
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			matrix[y][x] = "."
		}
	}
	return matrix
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func zapCoords(coords []Pos) []Pos {
	setLike := make(map[string]bool)
	for _, v := range coords {
		setLike[string(v.y)+string(v.x)] = true
	}
	var zapped []Pos
	for k := range setLike {
		zapped = append(zapped, Pos{y: int(k[0]), x: int(k[1])})
	}
	return zapped
}
