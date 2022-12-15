package main

import "fmt"

func run(){
	//declare a variable
	var name = "Breaking"
	// alter its value
	name = "Better Call Saul"

	var lastName string = "Bad"
	var age int = 55

fmt.Println("Hello %v",name)

fmt.Println("Hello World",name, lastName,age)

// city will be given a default value
 var city string
 
 city = "London"

 fmt.Println(city)

 var sport , subject , year = "foot ball" , "math" , 2022

 fmt.Println(sport,subject,year)

 var (	company = "MS"
		limit = 100
		speed = 166
 		)

	fmt.Println(company,limit,speed)		

	message := "We can avoid using var"
	fmt.Println(message)

	// we can reassgn with := ????

	//Logical operations

	var no1 = 55
	var no2 = 100
	var favAnimal = "horse"
	var ans = no1 < no2 && favAnimal == "horse"

	fmt.Println(ans);

}