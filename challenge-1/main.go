package main

import "fmt"

func main()  {

	var i = 21
	var j bool = true

	fmt.Println("challenge 1 ")

	fmt.Printf("%v \n", i) 
	fmt.Printf("%T \n", i) 
	fmt.Printf("%% \n") 
	fmt.Printf("%t \n\n", j)
	
	fmt.Printf("%b \n", i) 
	fmt.Printf("%c \n", '\u042F')
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U \n\n", 'Я')
	
	var k float64 = 123.456
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
	
	
}