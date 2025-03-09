package main

import (
	"iter"
	"math"
	"slices"
)

type day7 struct{}

func (d day7) Part1(lines iter.Seq[string]) any {
	return solveDay7(lines, func(moves int) int {
		return moves
	})
}

func (d day7) Part2(lines iter.Seq[string]) any {
	return solveDay7(lines, func(moves int) int {
		return moves * (moves + 1) / 2
	})
}

func solveDay7(lines iter.Seq[string], costFunc func(moves int) int) int {
	nums := parseDay7(lines)
	slices.Sort(nums)

	cost := math.MaxInt
	for n := nums[0]; n <= nums[len(nums)-1]; n++ {
		newCost := 0
		for _, num := range nums {
			newCost += costFunc(abs(num - n))
		}
		cost = min(newCost, cost)
	}
	return cost
}

func parseDay7(lines iter.Seq[string]) []int {
	next, stop := iter.Pull(lines)
	defer stop()
	l, _ := next()
	return parseNums([]byte(l))
}
