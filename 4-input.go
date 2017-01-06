package main

import "fmt"

func main() {
	age := 10
	name := "José"

	price := 14.45

	fmt.Printf("My age is %v\nMy name is %v\n",age,name)
	fmt.Printf("My price is %f\n\n",price);

	var name_full string

	fmt.Print("Enter your full name:")
	fmt.Scanf("%s\n",&name_full)
	fmt.Printf("Hi %s",name_full)

}