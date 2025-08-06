package main

import (
	"fmt"
)
// var email = "jansaytong_p@silpakorn.edu"

func main()  {
	//var name string = "Pannathorn"
	var age int = 20

	email := "jansaytong_p@silpakorn.edu"
	gpa := 4.00

	firstname, lastname := "Pannathorn", "Jansaytong"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age, email, gpa)
