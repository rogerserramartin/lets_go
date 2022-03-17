package main

import "fmt"

func main(){
	hello := sayHi()
	fmt.Println(hello)
	var result float32 = calcSum(45.34, 56.67)
	fmt.Println(result)
}

func sayHi() string {
	return "HI!!"
}

// in Go, unlike Java, name goes first and type goes second
func calcSum(a float32, b float32) float32 {
	return a + b
}