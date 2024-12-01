package solutions

import (
	"strconv"
	"strings"

	"github.com/julianchong00/aoc2024/utils"
)

func Day1(input *utils.Input) error {
	arr1 := make([]int64, 0)
	arr2 := make([]int64, 0)

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
		arr1 = append(arr1, int1)
		arr2 = append(arr2, int2)
	}

	diffTotal := 0
	// sorted1 :=
	return nil
}
