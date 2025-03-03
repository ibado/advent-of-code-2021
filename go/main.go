package main

import (
	"flag"
	"fmt"
	"iter"
	"time"
)

const AllDays = -1

type AoCDay interface {
	Part1(lines iter.Seq[string]) any
	Part2(lines iter.Seq[string]) any
}

var dayarr = []AoCDay{}

func runDay(dayNum int, day AoCDay) {
	runPart := func(day int, f func(iter.Seq[string]) any) (any, time.Duration) {
		start := time.Now()
		res := f(readLines(uint8(day)))
		return res, time.Since(start)
	}
	part1, time1 := runPart(dayNum, day.Part1)
	part2, time2 := runPart(dayNum, day.Part2)
	fmt.Printf("day%d part1: %-20v time: %v\n", dayNum, part1, time1)
	fmt.Printf("day%d part2: %-20v time: %v\n", dayNum, part2, time2)
	fmt.Println()
}

func main() {
	day := flag.Int("day", AllDays, "Run speciphic day")
	flag.Parse()

	if *day == AllDays {
		for i, day := range dayarr {
			runDay(i+1, day)
		}
	} else {
		assert(*day > 0)
		runDay(*day, dayarr[*day-1])
	}
}
