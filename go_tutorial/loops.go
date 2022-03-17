package main

import "fmt"
import "strconv"

/*
import (
	"fmt"
	"strconv"
)
*/

func main(){
	// slice is the same as array
	getNumbers()
	characters := []string{"Goku", "Jack", "Batman", "Queen", "Janna", "Pikachu"}
	characters = append(characters, "Chaos") // append doesn't modify the variable, it returns a new slice that assigns to the var characters
	fmt.Println(characters)
	for i:= 0 ; i < len(characters) ; i++ {
		fmt.Println(characters[i])
	}
	getEvens()
}

func getNumbers() {
	
	primes := [6]int{2, 3, 5, 7, 11, 13}
	for  i := 0 ; i < len(primes) ; i++{
		fmt.Println(strconv.Itoa(primes[i]))
	}	
}

func getEvens() {
	evens :=[3] int {2, 4, 6}
	for i, even := range evens{
		fmt.Println(i, even)
	}
}