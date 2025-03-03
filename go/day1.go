package main

import (
	"iter"
	"strconv"
)

type day1 struct{}

func (d day1) Part1(lines iter.Seq[string]) any {
	var prev, incr int
	for l := range lines {
		val, err := strconv.Atoi(l)
		Must(err)
		if prev != 0 && val > prev {
			incr++
		}
		prev = val

	}

	return incr
}

func (d day1) Part2(lines iter.Seq[string]) any {
	var measurements []int

	for l := range lines {
		val, err := strconv.Atoi(l)
		Must(err)
		measurements = append(measurements, val)
	}

	var prev, incr int
	for i := 0; i < len(measurements)-2; i++ {
		val := measurements[i] + measurements[i+1] + measurements[i+2]
		if prev != 0 && val > prev {
			incr++
		}
		prev = val
	}

	return incr
}
