package file

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntArray(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var shelf [][]int
	scanner := bufio.NewScanner(file)
	lineNo := 0
	var shelfItems []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			if lineNo > 0 {
				shelf = append(shelf, shelfItems)
			}
			shelfItems = nil
		} else {
			val, _ := strconv.Atoi(scanner.Text())
			shelfItems = append(shelfItems, val)
		}
		lineNo++
	}
	shelf = append(shelf, shelfItems)
	return shelf, scanner.Err()
}
