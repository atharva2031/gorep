package constants

import (
	"fmt"
)

const (
	KB = 1 << (10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main(){
	var filesize float64 = 40000000000
	fmt.Printf("The filesize is: %.2f GB", filesize/GB) 
	
}

