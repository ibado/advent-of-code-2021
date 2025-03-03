package main

import (
	"iter"
	"slices"
	"strconv"
	"strings"
)

type day3 struct{}

func (d day3) Part1(lines iter.Seq[string]) any {
	var arr [12]int
	for l := range lines {
		for i, c := range l {
			if c == '0' {
				arr[i] -= 1
			} else if c == '1' {
				arr[i] += 1
			} else {
				panic("invalid input")
			}
		}
	}
	var str strings.Builder
	for _, c := range arr {
		if c > 0 {
			str.WriteRune('1')
		} else {
			str.WriteRune('0')
		}
	}
	gamma, err := strconv.ParseUint(str.String(), 2, 64)
	Must(err)
	empsilon := 0b111111111111 ^ (gamma)

	return gamma * empsilon
}

func (d day3) Part2(lines iter.Seq[string]) any {
	nums := slices.Collect(lines)

	oxNums := slices.Clone(nums)
	for i := 0; i < len(nums[0]); i++ {
		bitCriteria := "1"
		bt := 0
		for _, oxNum := range oxNums {
			if oxNum[i:i+1] == "0" {
				bt--
			} else {
				bt++
			}
		}
		if bt < 0 {
			bitCriteria = "0"
		}
		var newOxNums []string
		for _, val := range oxNums {
			if val[i:i+1] == bitCriteria {
				newOxNums = append(newOxNums, val)
			}
		}
		oxNums = newOxNums
		if len(oxNums) == 1 {
			break
		}
	}
	co2Nums := slices.Clone(nums)
	for i := 0; i < len(nums[0]); i++ {
		bitCriteria := "0"
		bt := 0
		for _, co2Num := range co2Nums {
			if co2Num[i:i+1] == "0" {
				bt--
			} else {
				bt++
			}
		}
		if bt < 0 {
			bitCriteria = "1"
		}
		var newco2Nums []string
		for _, val := range co2Nums {
			if val[i:i+1] == bitCriteria {
				newco2Nums = append(newco2Nums, val)
			}
		}
		co2Nums = newco2Nums
		if len(co2Nums) == 1 {
			break
		}
	}
	oxRate, err := strconv.ParseInt(oxNums[0], 2, 64)
	Must(err)
	co2Rate, err := strconv.ParseInt(co2Nums[0], 2, 64)
	Must(err)
	return oxRate * co2Rate
}
