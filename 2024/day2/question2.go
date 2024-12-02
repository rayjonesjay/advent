package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Q2() {
	// read the file

	fd, _ := os.Open("question2_input.txt")

	// new scanner and scanning
	scanner := bufio.NewScanner(fd)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		// convert to list of integers
		list := ToList(words)

		if beta(list) || levelRemover(list) {
			count++
		}

	}
	fmt.Println(count)
}

func beta(arr []int) bool {
	isIncreasing := IsIncreasing(arr)
	for i := 0; i < len(arr)-1; i++ {
		if isIncreasing {
			// if its increasing difference will be postive
			diff := arr[i+1] - arr[i]
			// check if it is in range
			if !InRange(diff) {
				return false
			}
		} else {
			diff := arr[i] - arr[i+1]
			if !InRange(diff) {
				return false
			}
		}
	}
	return true
}

func levelRemover(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		tmp := make([]int, len(arr)-1)
		copy(tmp, arr[:i])
		copy(tmp[i:], arr[i+1:])
		if beta(tmp) {
			return true
		}
	}
	return false
}

func ToList(list []string) []int {
	arr := []int{}
	for _, l := range list {
		n, _ := strconv.Atoi(l)
		arr = append(arr, n)
	}
	return arr
}

func IsIncreasing(arr []int) bool {
	increaseCount := 0
	decreaseCount := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			increaseCount++
		} else {
			decreaseCount++
		}
	}
	if increaseCount > decreaseCount {
		return true
	} else if increaseCount == decreaseCount {
		return arr[0] < arr[1]
	}
	return false
}

func InRange(diff int) bool {
	return diff >= 1 && diff <= 3
}

func main() {
	Q2()
}
