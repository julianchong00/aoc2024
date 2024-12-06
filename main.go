package main

import (
	"fmt"
	"log"

	"github.com/julianchong00/aoc2024/solutions"
	"github.com/julianchong00/aoc2024/utils"
)

func main() {
	day := 3
	filename := fmt.Sprintf("inputs/%d.txt", day)

	input, err := utils.ReadInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = solutions.Day3(input)
	if err != nil {
		log.Fatal(err)
	}
}
