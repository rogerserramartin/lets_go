package main

import "fmt"


func main(){
	//types of variable declaration
	// both are static
	var car string = "seat panda"
	// go compiler decides that this is a string, but won't be reassigned into another type like Python or JS
	moto := "yamaha"

	decimals := 10.34

	one, two, three := 1, 2, 3 //like Python
	var s1, s2, s3 string = "one", "two", "three"

	var isTall bool = false

	// arrays
	/*
	var arrayNumbers [5]int
	arrayNumbers[0] = 1
	arrayNumbers[1] = 2
	arrayNumbers[2] = 3
	arrayNumbers[3] = 4
	arrayNumbers[4] = 5
	
	primes := [6]int{1,2,3,5,7,11}
	*/
	fmt.Println(car)
	fmt.Println(moto)
	fmt.Println(decimals)
	fmt.Println(isTall)
	fmt.Println(one, two, three)
	fmt.Println(s1, s2, s3)

	// basic types: bool, string, int, float64
}