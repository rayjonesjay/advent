package q2
import(
	//"fmt"
	"slices"
)

func Runner(l,w,h int) int {
	cube := l * w * h
	a := 2*l + 2*w
	b := 2*w + 2*h
	c := 2*l + 2*h
	res := []int{a,b,c}
	slices.Sort(res)
	return res[0] + cube
} 
