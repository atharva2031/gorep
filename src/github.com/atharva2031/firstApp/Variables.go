//exercise for datatypes in Go

package main
import (
	"fmt"
	"strconv"
)
var (
	k int = 45
	l string  = strconv.Itoa(k)
)

func main()  {
	var i int = 45
	var j float64
	j = 45.0
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", l, l)
}