package main

import (
	"fmt"
	"os"
	"petalv/aoc_2022/file"
	"sort"
	"sync"
)

type ShelfItem struct {
	Id       int
	Calories int
}

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadIntArray("input.txt")
	if part == "part2" {
		fmt.Println(countCalories(2, input))
	} else {
		fmt.Println(countCalories(1, input))
	}
}

func countCalories(part int, shelf [][]int) int {
	var shelfCount = len(shelf)
	mappedShelf := make(chan ShelfItem, shelfCount)
	reducedShelf := make(chan ShelfItem)
	var wg sync.WaitGroup
	wg.Add(shelfCount)

	for id, shelfItems := range shelf {
		go func(calories []int, id int) {
			defer wg.Done()
			mappedShelf <- Map(calories, id)
		}(shelfItems, id)
	}

	if part == 1 {
		go Reducer(mappedShelf, reducedShelf)
	} else {
		go Reducer2(mappedShelf, reducedShelf)
	}
	wg.Wait()

	close(mappedShelf)
	res := <-reducedShelf
	return res.Calories
}

func Map(calories []int, id int) ShelfItem {
	sum := addArray(calories)
	return ShelfItem{
		Id:       id,
		Calories: sum,
	}
}

func Reducer(shelves chan ShelfItem, reduced chan ShelfItem) {
	final := ShelfItem{Id: -1, Calories: 0}
	for shelf := range shelves {
		if shelf.Calories > final.Calories {
			final = shelf
		}
	}
	reduced <- final
}

func Reducer2(shelves chan ShelfItem, reduced chan ShelfItem) {
	var tops []int
	for shelf := range shelves {
		if len(tops) < 3 {
			tops = append(tops, shelf.Calories)
		} else {
			sort.Ints(tops)
			if shelf.Calories > tops[2] {
				tops[0] = tops[1]
				tops[1] = tops[2]
				tops[2] = shelf.Calories
			} else if shelf.Calories > tops[1] {
				tops[0] = tops[1]
				tops[1] = shelf.Calories
			} else if shelf.Calories > tops[0] {
				tops[0] = shelf.Calories
			}
		}
	}
	fmt.Printf("%v", tops)
	reduced <- ShelfItem{Id: -1, Calories: addArray(tops)}
}

func addArray(numbs []int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}
