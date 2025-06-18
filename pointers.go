package main

import "fmt"

// func main() {
// 	salary := 50000

// 	// Get a pointer to salary
// 	var salaryPointer *int = &salary

// 	fmt.Println("salary:", *salaryPointer)

// 	// Pass pointer to function
// 	taxableYears := getTaxableYears(salaryPointer)

// 	fmt.Println("Taxable salary years:", taxableYears)
// }

// // passing pointers to function
// func getTaxableYears(salary *int) int {
// 	return *salary / 10000
// }
func main() {
	

	totalUsers := 3000

	// make a pointer to total users 

	totalUsersPointer := &totalUsers




	activeUsersOnApp := getActiveUsers(totalUsersPointer)

	fmt.Println("active users on app", activeUsersOnApp)


}


func getActiveUsers (totalUsers *int) int {
	return *totalUsers - 150   // *totalUsers --> it is de referencing
}
