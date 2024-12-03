package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args[1:]
	fd := Reader(arg[0])
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	store := []string{}
	sum := 0

	pat := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pat)

	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		store = append(store, re.FindAllString(strings.TrimSpace(line), -1)...)
		sum2 += q2(line)
	}

	for _, s := range store {
		sum += calc(s)
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func q2(s string) int {
	s = strings.TrimSpace(s)
	pat2 := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	re2 := regexp.MustCompile(pat2)
	match2 := re2.FindAllStringSubmatch(s, -1)
	enabled := true // mul are enabled by defualt
	fmt.Println(s, match2)
	sum2 := 0
	for _, match := range match2 {
		fullMatch := match[0]

		if fullMatch == "do()" {
			fmt.Println(fullMatch)
			enabled = true
		} else if fullMatch == "don't()" {
			enabled = false
		} else if len(match) > 2 {
			if enabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				sum2 += x * y
				// fmt.Println(x, y, match, sum2)
			} else {
				fmt.Printf("\n%q->%q\n", match,s)
			}
		}
	}
	return sum2
}
func calc(s string) int {
	s = strings.TrimLeft(s, "mul")
	s = s[1 : len(s)-1]
	words := strings.Split(s, ",")
	a, _ := strconv.Atoi(words[0])
	b, _ := strconv.Atoi(words[1])
	return a * b
}
func Reader(filename string) *os.File {
	fd, _ := os.Open(filename)
	return fd
}
