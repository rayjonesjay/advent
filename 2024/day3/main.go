package main

import (
	"bufio"
	"os"
)

func main() {
	arg := os.Args[1:]
	fd := Reader(arg[0])
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		
	}
	defer fd.Close()
}

func Reader(filename string) *os.File {
	fd, _ := os.Open(filename)
	return fd
}
