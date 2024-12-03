package q1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Runner() {
	arg := os.Args[1:]
	if len(arg) != 1{
		fmt.Println("provide file input")
		return
	}
	fd := Reader(arg[0])
	defer fd.Close()
	
	scanner := bufio.NewScanner(fd)
	store := []string{}
	sum := 0

	pat := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pat)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		store = append(store, re.FindAllString(strings.TrimSpace(line), -1)...)
	}

	for _, s := range store {
		sum += Calc(s)
	}
	fmt.Println(sum)
}

func Calc(s string) int {
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
