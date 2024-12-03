package q2

import (
	"day3/q1"
	"fmt"
	"os"
	"regexp"
	"strings"
	"bufio"
)

func Runner() {
	do := "do()"
	dont := "don't()"

	enable := true // default is enabled
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println("please provide file input")
		return
	}

	fd := q1.Reader(arg[0])

	pat := `mul\(\d{1,3}\,\d{1,3}\)|don\'t\(\)|do\(\)`

	scanner := bufio.NewScanner(fd)

	re := regexp.MustCompile(pat)
	res := 0
	all := []string{}
	for scanner.Scan() {
		currLine := scanner.Text()
		currLine = strings.TrimSpace(currLine)
		all = append(all, re.FindAllString(currLine, -1)...)
	}

	for _, s := range all {
		if s == do {
			enable = true
		}else if s == dont{
			enable = false
		}else {
			if enable{
			res+=q1.Calc(s)
			}
		}
	}
	fmt.Println(res)
}
