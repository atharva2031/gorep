package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHQ
	canSeeFinancials
	canSeeSA
	canSeeNA
)

func main() {
	fmt.Println(byte(isAdmin))
	fmt.Println(byte(isHQ))
	fmt.Println(byte(canSeeFinancials))
	fmt.Println(byte(canSeeSA))
	fmt.Println(byte(canSeeNA))

	
}
