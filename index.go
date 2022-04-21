package main

import "fmt"

func hello() string {
	return "hello"
}

func main() {
	fmt.Print(hello())
	var a string = "hello"
	fmt.Print(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Print(primes[1:2])
}
