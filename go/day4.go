package main

import "iter"

type day4 struct{}

type BingoNum struct {
	val    int
	marked bool
}

func (d day4) Part1(lines iter.Seq[string]) any {
	nums, boards := parseDay4(lines)
	for _, n := range nums {
		for _, board := range boards {
			if markNum(board, n) {
				return score(board, n)
			}
		}
	}

	return 0
}

func (d day4) Part2(lines iter.Seq[string]) any {
	nums, boards := parseDay4(lines)
	var winnerIdx int
	var lastNum int
	winners := make(map[int]bool)
	for _, n := range nums {
		for boardIdx, board := range boards {
			if !winners[boardIdx] && markNum(board, n) {
				winnerIdx = boardIdx
				lastNum = n
				winners[boardIdx] = true
			}
		}
	}
	return score(boards[winnerIdx], lastNum)
}

func parseDay4(lines iter.Seq[string]) ([]int, [][][]BingoNum) {
	next, stop := iter.Pull(lines)
	defer stop()
	firsLine, _ := next()
	nums := parseNums([]byte(firsLine))

	var boards [][][]BingoNum
	var board [][]BingoNum
	next()
	for {
		line, ok := next()
		if !ok {
			break
		}
		if line == "" {
			boards = append(boards, board)
			board = nil
			continue
		}

		var numRow []BingoNum
		for _, n := range parseNums([]byte(line)) {
			numRow = append(numRow, BingoNum{n, false})
		}
		board = append(board, numRow)

	}

	return nums, boards
}

func markNum(board [][]BingoNum, num int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j].val == num {
				board[i][j].marked = true
				line := true
				for i := 0; i < 5; i++ {
					if !board[i][j].marked {
						line = false
					}
				}
				if line {
					return true
				}
				line = true
				for j := 0; j < 5; j++ {
					if !board[i][j].marked {
						line = false
					}
				}
				if line {
					return true
				}
				return false
			}
		}
	}
	return false
}

func score(board [][]BingoNum, num int) int {
	unmarkSum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board[i][j].marked {
				unmarkSum += board[i][j].val
			}
		}
	}
	return num * unmarkSum
}
