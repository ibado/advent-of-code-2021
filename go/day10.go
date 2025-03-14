package main

import (
	"iter"
	"slices"
)

type day10 struct{}

var openChars = []byte{
	'(', '[', '{', '<',
}

var closingChars = []byte{
	')', ']', '}', '>',
}

var points = []int{
	3, 57, 1197, 25137,
}

func (d day10) Part1(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	illegals := make(map[byte]int)
	for _, line := range mx {
		var stack []byte
		for _, c := range line {
			if idx := slices.Index(openChars, c); idx != -1 {
				stack = append(stack, c)
			} else if len(stack) == 0 {
				// incomplete
				break
			} else {
				toFind := openChars[slices.Index(closingChars, c)]
				poped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if poped != toFind {
					illegals[c]++
					break
				}
			}
		}
	}

	res := 0
	for i := 0; i < 4; i++ {
		res += illegals[closingChars[i]] * points[i]
	}

	return res
}

func (d day10) Part2(lines iter.Seq[string]) any {
	return 0
}
