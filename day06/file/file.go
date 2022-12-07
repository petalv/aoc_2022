package file

import (
	"bufio"
	"os"
)

func ReadStringArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var items []string
	for scanner.Scan() {
		val := scanner.Text()
		items = append(items, val)
	}
	return items, scanner.Err()
}
