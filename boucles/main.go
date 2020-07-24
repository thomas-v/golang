package main

import "fmt"

func main(){
	sum := 0;
	for i := 0; i < 10; i++ {
		sum += i;
	}
	fmt.Println(sum);

	sum = 0
	for sum < 5 {
		sum++
	}
	fmt.Println(sum)

	sum = 0
	for{
		sum++
		fmt.Println(sum)
	}
}

