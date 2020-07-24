package main

import "fmt"

type Personne struct{
	nom string
	prenom string
	age int
}

func main(){
	fmt.Println(Personne{"VIVI", "THOMAS", 30})
	personne1 := Personne{"VIVI", "THOMAS", 30}
	fmt.Printf("Le prenom de la personne est : %v\n", personne1.prenom)
	personne1.prenom = "ARNAUD"
	fmt.Printf("Le prenom de la personne est : %v\n", personne1.prenom)

	fmt.Println(*&personne1.age)
}
