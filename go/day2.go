package main

import (
	"iter"
	"strconv"
	"strings"
)

type day2 struct{}

func (d day2) Part1(lines iter.Seq[string]) any {
	var forward, depth int
	for l := range lines {
		command, num := parseDay2Line(l)
		switch command {
		case "forward":
			forward += num
		case "down":
			depth += num
		case "up":
			depth -= num
		default:
			panic("invalid command")
		}
	}
	return forward * depth
}

func (d day2) Part2(lines iter.Seq[string]) any {
	var forward, depth, aim int
	for l := range lines {
		command, num := parseDay2Line(l)
		switch command {
		case "forward":
			forward += num
			depth += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		default:
			panic("invalid command")
		}
	}
	return forward * depth
}

func parseDay2Line(l string) (string, int) {
	split := strings.Split(l, " ")
	num, err := strconv.Atoi(split[1])
	Must(err)
	return split[0], num
}
