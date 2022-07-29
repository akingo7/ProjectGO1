package main
import "fmt"
func getUserInfo() {
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&information.firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&information.lastName)
	fmt.Print("Please enter your age: ")
	fmt.Scan(&information.age)
	fmt.Print("Please enter your mail address: ")
	fmt.Scan(&information.mail)
}

func verifyUserInfoDL() bool {
	detailsLength := len(information.firstName) < 2 ||  len(information.lastName) < 2 || len(information.mail) > 4
	return detailsLength
}

func postUserInfo() {
	fmt.Print("\n######################\n")
	fmt.Printf("First Name: %v\nLast Name: %v\nAge: %v\nE-mail Address: %v\n", information.firstName, information.lastName, information.age, information.mail)
	fmt.Print("\n######################\n")

}
