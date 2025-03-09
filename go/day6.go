package main

import (
	"iter"
)

type day6 struct{}

func (d day6) Part1(lines iter.Seq[string]) any {
	next, stop := iter.Pull(lines)
	defer stop()
	line, _ := next()
	nums := parseNums([]byte(line))

	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	count := countFishes(80, nums)

	return count
}

func (d day6) Part2(lines iter.Seq[string]) any {
	next, stop := iter.Pull(lines)
	defer stop()
	line, _ := next()
	_ = parseNums([]byte(line))

	return 0 //countFishes(256, nums)
}

func countFishes(n int, nums []int) int {
	for i := 0; i < n; i++ {
		ln := len(nums)
		for j := 0; j < ln; j++ {
			num := nums[j]
			if num == 0 {
				nums[j] = 6
				nums = append(nums, 8)
			} else {
				nums[j]--
			}
		}
	}

	return len(nums)
}
