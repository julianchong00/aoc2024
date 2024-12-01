package utils

import (
	"bufio"
	"os"
)

type Input struct {
	Lines []string
}

func ReadInput(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	return &Input{
		Lines: fileLines,
	}, nil
}
