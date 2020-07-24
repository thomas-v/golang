package main

import (
	"fmt"
	"runtime"
)

func main(){
	avg := 11
	if avg > 10 {
		fmt.Println("Vous avez la moyenne")
	} else if avg < 10 {
		fmt.Println("Vous n'avez pas la moyenne")
	} else {
		fmt.Println("Vous avez pile la moyenne")
	}

        if avgBis:= 9; avgBis > 10 {
                fmt.Println("Vous avez la moyenne")
        } else if avgBis < 10 {
                fmt.Println("Vous n'avez pas la moyenne")
        } else {
                fmt.Println("Vous avez pile la moyenne")
        }

	switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("LINUX")
		case "windows":
			fmt.Println("WINDOWS")
		default:
			fmt.Println("AUTRE")
	}

	avg = 10
	switch {
		case avg > 10:
			fmt.Println("Vous avez la moyenne")
		case avg < 10:
			fmt.Println("Vous n'avez pas la moyenne")
		default:
			fmt.Println("Vous avez pile la moyenne")
	}
}
