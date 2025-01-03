package main

import (
	"os"
	"bufio"
	"fmt"
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
		fmt.Println("Error:", err)
		return
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	counter := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		list := strings.Fields(line)
		ln := toList(list)

		// Check if the report is safe or can be made safe
		if helper(ln) || canBeMadeSafe(ln) {
			counter++
		}
	}

	fmt.Println(counter)
}

