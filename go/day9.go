package main

import (
	"iter"
	"strconv"
)

type day9 struct{}

func (d day9) Part1(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	sum := 0
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[0]); j++ {
			if isLow(mx, i, j) {
				n, err := strconv.Atoi(string(mx[i][j]))
				Must(err)
				sum += 1 + n
			}
		}
	}
	return sum
}

func (d day9) Part2(lines iter.Seq[string]) any {
	return 0
}

func isLow(mx [][]byte, i, j int) bool {
	num := mx[i][j]
	if i > 0 && num >= mx[i-1][j] {
		return false
	}
	if j > 0 && num >= mx[i][j-1] {
		return false
	}
	if i > 0 && j > 0 && num >= mx[i-1][j-1] {
		return false
	}

	if i < len(mx)-1 && num >= mx[i+1][j] {
		return false
	}
	if j < len(mx[0])-1 && num >= mx[i][j+1] {
		return false
	}
	if i < len(mx)-1 && j < len(mx[0])-1 && num >= mx[i+1][j+1] {
		return false
	}

	if i > 0 && j < len(mx[0])-1 && num >= mx[i-1][j+1] {
		return false
	}

	if j > 0 && i < len(mx)-1 && num >= mx[i+1][j-1] {
		return false
	}
	return true
}
