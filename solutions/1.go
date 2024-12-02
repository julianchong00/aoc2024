package solutions

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/julianchong00/aoc2024/utils"
	"github.com/samber/lo"
)

func Day1(input *utils.Input) error {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	for _, line := range input.Lines {
		lineElems := strings.Split(line, "   ")
		int1, err := strconv.ParseInt(lineElems[0], 10, 64)
		if err != nil {
			return err
		}
		int2, err := strconv.ParseInt(lineElems[1], 10, 64)
		if err != nil {
			return err
		}
		arr1 = append(arr1, int(int1))
		arr2 = append(arr2, int(int2))
	}

	part1(arr1, arr2)
	part2(arr1, arr2)

	return nil
}

func part1(arr1, arr2 []int) {
	slices.SortFunc(arr1, func(a, b int) int {
		return a - b
	})
	slices.SortFunc(arr2, func(a, b int) int {
		return a - b
	})

	zip := lo.Zip2(arr1, arr2)
	diffTotal := 0
	for _, tup := range zip {
		diff := math.Abs(float64(tup.A) - float64(tup.B))
		diffTotal += int(diff)
	}

	fmt.Println("Total Distance between lists:", diffTotal)
}

func part2(arr1, arr2 []int) {
	totalSimilarityScore := 0

	occurrenceMap := map[int]int{}
	for _, num := range arr2 {
		if occurrence, found := occurrenceMap[num]; !found {
			occurrenceMap[num] = 1
		} else {
			occurrenceMap[num] = occurrence + 1
		}
	}

	for _, num := range arr1 {
		if occurrence, found := occurrenceMap[num]; !found {
			continue
		} else {
			totalSimilarityScore += num * occurrence
		}
	}

	fmt.Println("Total similarity score:", totalSimilarityScore)
}
