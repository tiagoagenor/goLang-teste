package main

import "fmt"
import "strconv"

func main() {
	age := 10
	age_str := strconv.Itoa(age) 

	age2 := "10"
	age2_int,_ := strconv.Atoi(age2)


	fmt.Println("My age is " + age_str)
	fmt.Println(age2_int + 10)
}