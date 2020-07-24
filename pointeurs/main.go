package main

import (
	"fmt"
)

func main(){
	var a int = 10
	var pa *int = &a
	fmt.Printf("L'adresse memoire de la variable a est : %p\n", pa)
	fmt.Printf("L'adresse memoire de la variable a est : %v\n", &a)
	fmt.Printf("La valeur de la variable a est : %v\n", a)
	fmt.Printf("La valeur de la variable a via le pointeur pa est : %v\n", *pa)
	fmt.Printf("La valeur de la variable a est : %v\n", *&a)
}
