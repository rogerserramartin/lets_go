package main // the name != to the folder where it belongs to
// main is used to make an executable package instead of a reusable one
// main.exe runs the func main() when executed

import "fmt"  // https://pkg.go.dev/fmt ->  fmt implements formatted I/O with functions, similar to C and C++
// fmt stands for format

func main(){
	fmt.Println("Hello World!")
}