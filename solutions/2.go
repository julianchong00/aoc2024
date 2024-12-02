package solutions

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/julianchong00/aoc2024/utils"
	"github.com/samber/lo"
)

func Day2(input *utils.Input) error {
	reports := make([][]int, 0)

	for _, line := range input.Lines {
		lineElems := strings.Split(line, " ")
		reports = append(reports, lo.Map(lineElems, func(valStr string, _ int) int {
			val, err := strconv.ParseInt(valStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			return int(val)
		}))
	}

	day2part1(reports)
	day2part2(reports)

	return nil
}

func day2part1(reports [][]int) {
	totalSafeReports := 0

	for _, report := range reports {
		isSafe := levelChecker(report)
		if isSafe {
			totalSafeReports++
		}
	}

	fmt.Println("Total number of safe reports:", totalSafeReports)
}

func day2part2(reports [][]int) {
	totalSafeReports := 0

	for _, report := range reports {
		isSafe := levelChecker(report)
		if isSafe {
			totalSafeReports++
		} else {
			for i := 1; i <= len(report); i++ {
				newReport := slices.Concat(report[0:i-1], report[i:])
				newIsSafe := levelChecker(newReport)
				if newIsSafe {
					totalSafeReports++
					break
				}
			}
		}
	}

	fmt.Println("Total number of safe reports with Problem Dampener:", totalSafeReports)
}

func levelChecker(report []int) bool {
	isAscending := lo.IsSorted(report)
	isDescending := true

	for i := 1; i < len(report); i++ {
		if report[i-1] <= report[i] {
			isDescending = false
		}
	}

	safeDiff := true
	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i-1]) - float64(report[i]))
		if int(diff) < 1 || int(diff) > 3 {
			safeDiff = false
		}
	}

	isSafe := (isAscending || isDescending) && safeDiff
	return isSafe
}
