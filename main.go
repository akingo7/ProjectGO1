package main

type info struct {
	firstName string
	lastName  string
	age       uint
	mail      string
}

var information info

func main() {
	for i := 1; i <= 20; i++ {
		getUserInfo()
		if verifyUserInfoDL() {
			postUserInfo()
		} else {
			verifyUserInfoFB()
		}
	}
}
