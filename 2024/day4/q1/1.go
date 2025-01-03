package q1

import (
	"fmt"
	"bufio"
	"os"
)

func Run(){
	fmt.Println("helo from q1")
	arg := os.Arg[1:]

	if len(arg) != 1 {
		fmt.Println("Usage: go run . input_file.txt")
		return
	}

	fd := Reader(filename)
	defer fd.Close()
	
	scanner := bufio.NewScanner(fd)

	grid := [][]rune{}
	for scanner.Scan(){
		line := scanner.Text()
		arr := []rune{}
		arr = append(arr,line...)
	}

}

func Reader(fname string) (*os.File) {
	fd, _ := os.Open(fname)
	return fd
}