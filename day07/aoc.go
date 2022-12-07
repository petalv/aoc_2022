package main

import (
	"fmt"
	"os"
	"petalv/aoc_2022/file"
	"sort"
	"strconv"
	"strings"
)

const (
	UNKNOWN = -1
	PUSH    = 1
	POP     = 2
	LIST    = 3
)

type Command struct {
	cmd     int
	mutator string
	output  []string
}
type Node struct {
	name     string
	parent   *Node
	children []*Node
	files    []File
	size     int
}

type File struct {
	size int
	name string
}

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStringArray("input.txt")
	part1Result, part2Result := handle(input)
	if part == "part2" {
		println(part2Result)
	} else {
		println(part1Result)
	}
}

func handle(input []string) (int, int) {
	FS_TOTAL := 70000000
	FS_REQUIRED := 30000000

	var root = &Node{name: "/"}
	var commands []*Command
	for _, console := range input {
		if strings.HasPrefix(console, "$") {
			command, mutator := parseCommand(console)
			commands = append(commands, &Command{cmd: command, mutator: mutator})
		} else {
			var cmd = commands[len(commands)-1]
			cmd.output = append(cmd.output, console)
		}
	}
	var pwd = &Node{children: []*Node{root}}
	for _, c := range commands {
		// fmt.Printf("Dir %s\n", pwd.name)
		// fmt.Printf("[%d] command: %d %s %s\n", i, c.cmd, c.mutator, strings.Join(c.output, ";"))
		if c.cmd == PUSH {
			if c.mutator == "/" {
				pwd = root
			} else {
				for _, child := range pwd.children {
					if child.name == c.mutator {
						pwd = child
					}
				}
			}
		} else if c.cmd == POP {
			pwd = pwd.parent
		} else if c.cmd == LIST {
			for _, file := range c.output {
				if strings.HasPrefix(file, "dir") {
					pwd.children = append(pwd.children, &Node{name: file[4:], parent: pwd})
				} else {
					fSplit := strings.Split(file, " ")
					fSize, _ := strconv.Atoi(fSplit[0])
					pwd.files = append(pwd.files, File{name: fSplit[1], size: fSize})
				}
			}
		}
	}
	sumSize(root)
	freespace := FS_TOTAL - root.size
	neededspace := FS_REQUIRED - freespace
	part1 := findSize(root, 100000)
	part2 := findSizeMin(root, neededspace)
	var part2Sizes []int
	for _, n := range part2 {
		part2Sizes = append(part2Sizes, n.size)
	}
	sort.Ints(part2Sizes)

	return part1, part2Sizes[0]
}

func findSize(node *Node, maxSize int) int {
	sum := 0
	for _, c := range node.children {
		if c.size <= maxSize {
			sum += sumDir(c)
		} else {
			sum += findSize(c, maxSize)
		}
	}
	return sum
}

func findSizeMin(node *Node, minSize int) []*Node {
	var targets []*Node
	if node.size >= minSize {
		targets = append(targets, node)
		for _, c := range node.children {
			targets = append(targets, findSizeMin(c, minSize)...)
		}
	}

	return targets
}

func sumSize(node *Node) int {
	for _, f := range node.files {
		node.size += f.size
	}
	for _, c := range node.children {
		childrenSize := sumSize(c)
		node.size += childrenSize
	}
	return node.size
}

func sumDir(node *Node) int {
	sum := 0
	for _, c := range node.children {
		sum += sumDir(c)
	}
	return node.size + sum
}

func tree(node *Node, depth int) {
	spacing := strings.Repeat(" ", depth+2)
	fmt.Printf("%s- %s (dir)\n", spacing, node.name)
	for i, c := range node.children {
		tree(c, depth+(i))
	}
	for _, f := range node.files {
		fmt.Printf("%s- %s (file, size=%d)\n", spacing, f.name, f.size)
	}
}

func parseCommand(command string) (int, string) {
	if strings.HasPrefix(command, "$ cd ") {
		if command[5:] == ".." {
			return POP, ""
		}
		return PUSH, command[5:]
	}
	if strings.HasPrefix(command, "$ ls") {
		return LIST, ""
	}
	return UNKNOWN, ""
}
