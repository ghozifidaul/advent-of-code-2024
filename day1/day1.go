package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func getInput() ([]int, []int) {
	file, err := os.Open("day1-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // or any other split function
	var locationIds []int
	for scanner.Scan() {
		word := scanner.Text()
		locationid, err := strconv.Atoi(word)

		if err != nil {
			fmt.Println("error during conversion")
		}

		locationIds = append(locationIds, locationid)
	}

	var leftArr []int
	var rightArr []int
	for i := 0; i < len(locationIds); i++ {
		if i%2 == 0 {
			leftArr = append(leftArr, locationIds[i])
		} else {
			rightArr = append(rightArr, locationIds[i])
		}
	}

	leftSortedArr := bubbleSort(leftArr)
	rightSortedArr := bubbleSort(rightArr)

	return leftSortedArr, rightSortedArr
}

func countDistance(leftSortedArr []int, rightSortedArr []int) {
	totalDistance := 0
	for i := 0; i < len(leftSortedArr); i++ {
		distance := 0
		if leftSortedArr[i] > rightSortedArr[i] {
			distance = leftSortedArr[i] - rightSortedArr[i]
		} else {
			distance = rightSortedArr[i] - leftSortedArr[i]
		}
		totalDistance = totalDistance + distance
	}

	fmt.Println("Total distance:", totalDistance)
}

func countSimilarity(leftSortedArr []int, rightSortedArr []int) {
	totalSimilarity := 0
	for i := 0; i < len(leftSortedArr); i++ {
		n := leftSortedArr[i]
		totalAppearance := 0

		for j := 0; j < len(rightSortedArr); j++ {
			if n == rightSortedArr[j] {
				totalAppearance = totalAppearance + 1
			}
		}
		totalSimilarity = totalSimilarity + (totalAppearance * n)
	}
	fmt.Println("Total Similarity: ", totalSimilarity)
}

func main() {
	leftSortedArr, rightSortedArr := getInput()
	// countDistance(leftSortedArr, rightSortedArr)
	countSimilarity(leftSortedArr, rightSortedArr)
}
