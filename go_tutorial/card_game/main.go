package main // this will be an executable package

import (
	"fmt"
)


type deck[]string

func main(){ // when we have a package called main, we must also declare a main func
	cards := deck{"Ace of Spades", "Six of Hearts"}
	fmt.Println(cards)

}

func newCard() string{
	return "Ace of Spades"
}