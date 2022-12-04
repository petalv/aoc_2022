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
	var rounds []string
	for scanner.Scan() {
		val := scanner.Text()
		rounds = append(rounds, val)
		// fmt.Println(val)
	}
	return rounds, scanner.Err()
}
