package main

import "fmt"

func main() {
	fmt.Println("challenge - 3")
	str := "selamat malam"

	data := map[string]int{}
	
	for _, v := range str {
		greet := string(v)
		fmt.Println(greet)
		data[greet] += 1
	}
	fmt.Println(data)
}