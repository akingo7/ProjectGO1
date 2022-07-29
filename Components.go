package main
import "fmt"
func getUserInfo() {
	fmt.Print("\nPlease enter your first name: ")
	fmt.Scan(&information.firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&information.lastName)
	fmt.Print("Please enter your age: ")
	fmt.Scan(&information.age)
	fmt.Print("Please enter your mail address: ")
	fmt.Scan(&information.mail)
}

var nameGreaterThanNu int = 2
var firstNameCondition bool = len(information.firstName) > nameGreaterThanNu
var lastNameCondition bool = len(information.lastName) > nameGreaterThanNu
var mailCondition bool = len(information.mail) > 4

func verifyUserInfoDL() bool {
	detailsLength := firstNameCondition && lastNameCondition && mailCondition
	return detailsLength
}

func postUserInfo() {
	fmt.Print("\n######################\n")
	fmt.Printf("First Name: %v\nLast Name: %v\nAge: %v\nE-mail Address: %v\n", information.firstName, information.lastName, information.age, information.mail)
	fmt.Print("\n######################\n")
}

func verifyUserInfoFB() {
	if !firstNameCondition {
		fmt.Print("\nThe first name you entered is invalid\n")
	} 
	if !lastNameCondition {
		fmt.Print("The last name you entered is invalid\n")
	}
	if !mailCondition {
		fmt.Print("You have entered an invalid mail\n")
	}
}
