package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	figures := figure{"Ironbreaker", "Gyrocopter", "Quarreler"}
	for i, fig := range figures {
		fmt.Println(i, fig)

		palabra := "coche"
		fmt.Println(len(palabra))
		fmt.Println(len(figures))

	}

	lidia := person{"Lidia", 25}
	fmt.Println(lidia.name)

}
