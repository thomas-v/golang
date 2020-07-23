package main

import (
	"fmt"
)

/*
	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias pour uint8

	rune // alias pour int32
	     // représente un "code point" Unicode

	float32 float64

	complex64 complex128	
*/

var varInt int
var varString string
var varBool, varBoolBis bool

var name, nameBis string = "Neo", "Morpheus"
var movie = "Matrix"

func main() {

	var varIntBis int
	fmt.Printf("Le résultat de la variable 'varInt' est : %d\n", varInt)
	fmt.Printf("Le résultat de la variable 'varString est : %v'\n", varString)
	if varString == "" {
		fmt.Printf("La variable 'varString' est vide\n")
	} else {
		fmt.Printf("La variable 'varString' est remplie\n")
	}
	fmt.Printf("Le résultat de la variable 'varBool' est : %v\n", varBool)
	fmt.Printf("Le résultat de la variable 'varBoolBis' est : %v\n", varBoolBis)
	fmt.Printf("Le résultat de la variable 'varIntBis' est : %d\n", varIntBis)
	
	fmt.Printf("Le résultat de la variable 'name' est : %v\n", name)
	fmt.Printf("Le résultat de la variable 'nameBis' est : %v\n", nameBis)

	fmt.Printf("Le résultat de la variable 'movie' est : %v et son type est : %T\n", movie, movie)

	diplomaName, diplomaType, diplomaNote := "SIO", "BTS", 14
	fmt.Printf("Le résultat de la variable 'diplomaName' est : %v et son type est : %T\n", diplomaName, diplomaName)
	fmt.Printf("Le résultat de la variable 'diplomaType' est : %v et son type est : %T\n", diplomaType, diplomaType)
	fmt.Printf("Le résultat de la variable 'diplomaNote' est : %v et son type est : %T\n", diplomaNote, diplomaNote)
}
