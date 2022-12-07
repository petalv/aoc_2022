package main

import (
	"fmt"
	"os"
	"petalv/aoc_2022/file"
)

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStringArray("input.txt")
	if part == "part2" {
		result := decoder(input, 14)
		fmt.Printf("%d\n", result)
	} else {
		result := decoder(input, 4)
		fmt.Printf("%d\n", result)
	}
}

func decoder(input []string, length int) int {
	msg := input[0]
	for i := 3; i < len(msg)-length; i++ {
		var strMap = make(map[string]bool)
		subMsg := msg[i : i+length]
		for _, c := range subMsg {
			if strMap[string(c)] == true {
				break
			}
			strMap[string(c)] = true
		}
		if len(strMap) == length {
			return i + length
		}
	}
	return -1
}
