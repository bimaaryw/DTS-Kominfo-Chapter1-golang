package main

import "fmt"

func main() {

	fmt.Println("challenge 2 ")

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j <= 10; j++ {

	if j == 5 {
	uni := []rune{'С','А','Ш','А','Р','В','О'}

	for k := 0; k < len(uni); k++ {
		fmt.Printf("Characters %U '%c' starts at byte position %d \n", rune(uni[k]),rune(uni[k]), k*2)
	}

	} else {
		fmt.Println("Nilai j = ", j)
	}

	}

	
}