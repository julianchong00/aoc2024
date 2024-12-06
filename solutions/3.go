package solutions

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/julianchong00/aoc2024/utils"
)

var (
	mulRegex  = regexp.MustCompile(`(?:mul)\(\d{1,3},\d{1,3}\)`)
	doRegex   = regexp.MustCompile(`(?:do)\(\)`)
	dontRegex = regexp.MustCompile(`(?:don't)\(\)`)
)

func Day3(input *utils.Input) error {
	err := day3part1(input.Lines)
	if err != nil {
		return err
	}
	err = day3part2(input.Lines)
	if err != nil {
		return err
	}

	return nil
}

func day3part1(lines []string) error {
	matches := make([]string, 0)
	for _, line := range lines {
		strings := mulRegex.FindAllString(line, -1)
		matches = append(matches, strings...)
	}

	pairs := make([][]int, 0)
	for _, match := range matches {
		pair := []int{}
		match, _ = strings.CutPrefix(match, "mul(")
		match, _ = strings.CutSuffix(match, ")")
		numStr := strings.Split(match, ",")
		firstNum, err := strconv.ParseInt(numStr[0], 10, 64)
		if err != nil {
			return err
		}
		secondNum, err := strconv.ParseInt(numStr[1], 10, 64)
		if err != nil {
			return err
		}
		pair = append(pair, int(firstNum))
		pair = append(pair, int(secondNum))
		pairs = append(pairs, pair)
	}

	total := 0
	for _, pair := range pairs {
		total += pair[0] * pair[1]
	}

	fmt.Println("Part 1 total:", total)
	return nil
}

type match struct {
	idx        int
	value      int
	enableMul  bool
	disableMul bool
}

func day3part2(lines []string) error {
	total := 0
	mulEnabled := true
	for _, line := range lines {
		matches := make([]match, 0)
		mulIdxMatches := make([][]int, 0)
		doIdxMatches := make([][]int, 0)
		dontIdxMatches := make([][]int, 0)
		mulIdx := mulRegex.FindAllStringIndex(line, -1)
		mulIdxMatches = append(mulIdxMatches, mulIdx...)

		doIdx := doRegex.FindAllStringIndex(line, -1)
		doIdxMatches = append(doIdxMatches, doIdx...)

		dontIdx := dontRegex.FindAllStringIndex(line, -1)
		dontIdxMatches = append(dontIdxMatches, dontIdx...)

		for _, mulMatch := range mulIdxMatches {
			mulStr := line[mulMatch[0]:mulMatch[1]]
			mulStr, _ = strings.CutPrefix(mulStr, "mul(")
			mulStr, _ = strings.CutSuffix(mulStr, ")")
			numStr := strings.Split(mulStr, ",")
			firstNum, err := strconv.ParseInt(numStr[0], 10, 64)
			if err != nil {
				return err
			}
			secondNum, err := strconv.ParseInt(numStr[1], 10, 64)
			if err != nil {
				return err
			}

			matches = append(matches, match{
				idx:        mulMatch[0],
				enableMul:  false,
				disableMul: false,
				value:      int(firstNum) * int(secondNum),
			})
		}

		for _, doMatch := range doIdxMatches {
			matches = append(matches, match{
				idx:        doMatch[0],
				enableMul:  true,
				disableMul: false,
				value:      0,
			})
		}

		for _, dontMatch := range dontIdxMatches {
			matches = append(matches, match{
				idx:        dontMatch[0],
				enableMul:  false,
				disableMul: true,
				value:      0,
			})
		}

		sort.SliceStable(matches, func(i, j int) bool {
			return matches[i].idx < matches[j].idx
		})

		for _, match := range matches {
			if match.disableMul {
				mulEnabled = false
			}
			if match.enableMul {
				mulEnabled = true
			}

			if mulEnabled {
				total += match.value
			}
		}
	}

	fmt.Println("Part 2 total:", total)
	return nil
}
