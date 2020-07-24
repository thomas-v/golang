package main

import (
	"fmt"
)

func main(){
	fmt.Println(add(5,5))
	fmt.Println(getNames())
	fmt.Println(split(17))

	defer fmt.Println("World !")
	fmt.Println("Hello")
}

func add(x, y int) int{
	return x + y;
}

func getNames() (string, string, string) {
	return "Thomas", "Andre", "Arnaud"
}

func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}
