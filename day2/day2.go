package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData() [][]int {
	file, err := os.Open("day2-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var levels [][]int
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		var converted []int
		for i := 0; i < len(words); i++ {
			convert, err := strconv.Atoi(words[i])

			if err != nil {
				fmt.Println("Error during conversion")
			}

			converted = append(converted, convert)
		}

		levels = append(levels, converted)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return levels
}

func isIncreasingSafely(levels []int) bool {
	if len(levels) <= 1 {
		return true
	} else {
		a := levels[0]
		b := levels[1]

		// if a is less than be it's increasing
		if a < b {
			// check the difference
			diff := b - a

			// if the diff is more than 3 then return false
			if diff > 3 {
				return false
			}

			return isIncreasingSafely(levels[1:])
		} else {
			return false
		}
	}
}

func isDecreasingSafely(levels []int) bool {
	if len(levels) <= 1 {
		return true
	} else {
		a := levels[0]
		b := levels[1]

		// if a is less than be it's increasing
		if a > b {
			// check the difference
			diff := a - b

			// if the diff is more than 3 then return false
			if diff > 3 {
				return false
			}

			return isDecreasingSafely(levels[1:])
		} else {
			return false
		}
	}
}

func findProblematicLevel(levels []int) int {
	if levels[0] < levels[1] {
		for i := 0; i < len(levels); i++ {
			a := levels[i]
			b := levels[i+1]

			if a < b {
				diff := b - a

				if diff > 3 {
					if i != 0 {
						if levels[i+1]-levels[i-1] > 3 {
							return i + 1
						} else {
							return i
						}
					} else {
						return i
					}
				}
			} else {
				return i + 1
			}
		}
	} else if levels[0] > levels[1] {
		for i := 0; i < len(levels); i++ {
			a := levels[i]
			b := levels[i+1]

			if a > b {
				diff := a - b

				if diff > 3 {
					if i != 0 {
						if levels[i-1]-levels[i+1] > 3 {
							return i+1
						} else {
							return i
						}
					} else {
						return i
					}
				}
			} else {
				return i + 1
			}
		}
	} else {
		return 0
	}
	return -1
}

func removeProblematicLevel(levels []int, problematicIndex int) []int {
	var newArr []int
	for i := 0; i < len(levels); i++ {
		if i != problematicIndex {
			newArr = append(newArr, levels[i])
		}
	}

	return newArr
}

func problemDampener(levels []int) bool {
	// 1. check safety
	isSafe := checkSafety(levels)
	fmt.Println("Levels: ", levels)

	if !isSafe {
		problemIdx := findProblematicLevel(levels)
		fmt.Println("Promblem index: ", problemIdx)
		if problemIdx == -1 {
			return true
		}
		newLevels := removeProblematicLevel(levels, problemIdx)
		fmt.Println("New Levels: ", newLevels)
		isNewLevelsSafe := checkSafety(newLevels)
		fmt.Println("is the level save? ", isNewLevelsSafe)
		fmt.Println("===================")
		return isNewLevelsSafe
	}

	fmt.Println("is the level save? ", isSafe)
	fmt.Println("===================")
	return isSafe
}

func checkSafety(levels []int) bool {
	if levels[0] < levels[1] {
		res := isIncreasingSafely(levels)
		return res
	} else {
		res := isDecreasingSafely(levels)
		return res
	}
}

func main() {
	fmt.Println(problemDampener([]int{1,5,9,9,10}))
	// levels := getData()
	// safeCount := 0
	// for i := 0; i < len(levels); i++ {
	// 	if problemDampener(levels[i]) {
	// 		safeCount++
	// 	}
	// }
	// fmt.Println(safeCount)
}
