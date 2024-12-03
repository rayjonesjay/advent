package q1

import(
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"slices"
	"day2/q2"
)

func Run(){
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("go run . <inputfile.txt>")
		return
	}
	
	fd := Reader(arg[0])
	scanner := Scanner(fd)
	res := 0
	feet := 0
	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(line,"x")
		l,w,h := list[0],list[1],list[2]
		ll , _ := strconv.Atoi(l)
		ww , _ := strconv.Atoi(w)
		hh , _ := strconv.Atoi(h)
		feet += q2.Runner(ll,ww,hh)
		res = res + (2*ll*ww) + (2*ww*hh) + (2*hh*ll) + FindSmallArea(ll,ww,hh)
	}
	fmt.Println("square feet",res)
	fmt.Println("feet",feet)
}

func Scanner(fd *os.File) *bufio.Scanner {
	return bufio.NewScanner(fd)
}

func Reader(filename string) (*os.File) {
	fd, _ := os.Open(filename)
	return fd
}

func FindSmallArea(l,w,h int) int {
	a := l * w 
	b := w * h
	c := h * l
	res := []int{a,b,c}
	slices.Sort(res)
	return res[0]
}
