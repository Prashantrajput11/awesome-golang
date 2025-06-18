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

	outputUserDetails(appUser)
}

func outputUserDetails(u User){
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}