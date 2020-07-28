package main

import (
	"fmt"
)

func main(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"Thomas",
		"David",
		"Noa",
		"Pierre",
	}
	b := names[0:3]
	fmt.Println(names)
	fmt.Println(b)

	c := []int{1,2,3,4,5}
	fmt.Println(c)

	d:= []struct{
		a int
		b string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}
	fmt.Println(d)
}
