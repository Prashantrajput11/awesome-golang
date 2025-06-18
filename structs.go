package main

// import statements
import (
	"fmt"

)


// creating / defining  struct
type User struct{
firstName string 
lastName string 
birthdate string
}


// attaching func to a struct : basically called method
// the arg is called receiver | u User 
func ( u User )outputUserDetails(){
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// main func 
func main() {

	
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var appUser User 

	// Instantiationg a struct
	appUser =  User{
		firstName : userFirstName,
		lastName : userLastName,
		birthdate:  userBirthDate,
	}

	// fmt.Println(firstName, lastName, birthdate)

	appUser.outputUserDetails()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}