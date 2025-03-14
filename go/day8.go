package main

import (
	"iter"
	"slices"
	"strings"
)

type day8 struct{}

func (d day8) Part1(lines iter.Seq[string]) any {
	count := 0
	uniqs := []int{2, 3, 4, 7}
	for l := range lines {
		patterns, vals := parseDay8Line(l)
		for _, ptrn := range patterns {
			ptrnlen := len(ptrn)
			if slices.Contains(uniqs, ptrnlen) {
				for _, val := range vals {
					if len(val) == ptrnlen {
						count++
					}
				}
			}
		}
	}
	return count
}

func (d day8) Part2(lines iter.Seq[string]) any {
	return 0
}

func parseDay8Line(l string) ([10]string, [4]string) {
	split := strings.SplitN(l, "|", 2)
	patterns := strings.SplitN(strings.TrimSpace(split[0]), " ", 10)
	vals := strings.SplitN(strings.TrimSpace(split[1]), " ", 4)

	return [10]string(patterns), [4]string(vals)
}
