package main

import (
	"iter"
)

type day5 struct{}

func (d day5) Part1(lines iter.Seq[string]) any {
	return solveDay5(lines, false)
}

func (d day5) Part2(lines iter.Seq[string]) any {
	return solveDay5(lines, true)
}

func solveDay5(lines iter.Seq[string], useDiagnals bool) int {
	visited := make(map[Point]int)
	for l := range lines {
		nums := parseNums([]byte(l))
		x1, y1, x2, y2 := nums[0], nums[1], nums[2], nums[3]

		if x1 == x2 && y1 != y2 {
			y0 := min(y1, y2)
			for yi := y0; yi <= y0+abs(y1-y2); yi++ {
				p := Point{x1, yi}
				visited[p]++
			}
		} else if y1 == y2 && x1 != x2 {
			x0 := min(x1, x2)
			for xi := x0; xi <= x0+abs(x1-x2); xi++ {
				p := Point{xi, y1}
				visited[p]++
			}
		} else {
			if useDiagnals {
				x0 := min(x1, x2)
				m := (y1 - y2) / (x1 - x2)
				for xi := x0; xi <= x0+abs(x1-x2); xi++ {
					yi := m*(xi-x1) + y1
					visited[Point{xi, yi}]++
				}
			}
		}
	}
	res := 0
	for _, v := range visited {
		if v >= 2 {
			res++
		}

	}
	return res
}
