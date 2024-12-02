package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)
func main(){

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("go run main.go <input_file.txt>")
		return
	}

	fd , err := os.Open(args[0])

	if err != nil {
		fmt.Println("err",err)
		return
	}


	scanner := bufio.NewScanner(fd)
	counter := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		list := strings.Fields(line)
		ln := toList(list)
		ln = tolerate(ln)
		if helper(ln){
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

func toList(list []string) []int{
	arr := []int{}
	for _, l := range list {
		n, _ := strconv.Atoi(l)
		arr = append(arr,n)
	}
	return arr
}

// for question1
func helper(list []int) bool {
	arr := list
	isIncreasing := isIncreasing(arr)
	for i := 0; i < len(arr)-1; i++{
		if isIncreasing {
			diff := arr[i+1] - arr[i]
			if !(diff > 0 && inRange(diff)) {
				return false
			}
		}else {
			diff := arr[i] - arr[i+1]
			if !(diff > 0 && inRange(diff)){
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
	for i := 0; i < len(arr)-1; i++{
		if arr[i] < arr[i+1] {
			increaseCount++
		}else{
			decreaseCount++
		}
	}
	if increaseCount > decreaseCount {
		return true
	}
	return false
}

func tolerate(list []int) []int {
	isIncreasing := isIncreasing(list)
	fmt.Println("before",list)

	for i := 0; i < len(list)-1; i++ {

		if isIncreasing {
			diff := list[i+1] - list[i]
			if diff <= 0 {
				list = pop(list,i)
				fmt.Println("after",list)
				// fmt.Println("level",i)
				return list
			}
		}else {
			diff := list[i+1] - list[i]
			if diff >= 0 {
				list = pop(list,i)
				fmt.Println("after",list)
				// fmt.Println("level",i)
				return list
			}
		}
	}

	return list
}

func pop(list []int, i int) []int {
	list = append(list[:i], list[i+1:]...)
	return list
}