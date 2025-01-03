package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("go run main.go <input_file.txt>")
		return
	}

	fd, err := os.Open(args[0])

	if err != nil {
		fmt.Println("err", err)
		return
	}

	scanner := bufio.NewScanner(fd)
	counter := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		list := strings.Fields(line)
		ln := toList(list)
		if helper(ln) {
			counter++
		}

	}
	fmt.Println(counter)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func toList(list []string) []int {
	arr := []int{}
	for _, l := range list {
		n, _ := strconv.Atoi(l)
		arr = append(arr, n)
	}
	return arr
}

// for question1
func helper(list []int) bool {
	arr := list
	isIncreasing := isIncreasing(arr)
	for i := 0; i < len(arr)-1; i++ {
		if isIncreasing {
			diff := arr[i+1] - arr[i]
			if !(diff > 0 && inRange(diff)) {
				return false
			}
		} else {
			diff := arr[i] - arr[i+1]
			if !(diff > 0 && inRange(diff)) {
				return false
			}
		}
	}
	return true
}

func inRange(diff int) bool {
	return diff >= 1 && diff <= 3
}

func isIncreasing(arr []int) bool {
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

func pop(list []int, i int) []int {
	list = append(list[:i], list[i+1:]...)
	return list
}
