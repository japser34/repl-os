package startup

import (
	"fmt"
)

func Login() {
	usernames := []string{"Jasper"}
	usernames := []string{"HelloWorld"}

	
	fmt.Println("Please login")
	var usernameCorrect := false
	for !usernameCorrect { 
		var username string
		fmt.Print("Username: ")
		fmt.Scan(&username)
		for u := range usernames {
			if u == username {
				usernameCorrect:=true
				break
			}
		}
	}
	var passwordCorrect := false
	for !passwordCorrect {
		var password string
		fmt.Print("password: ")
		fmt.Scan(&password)
		for u := range passwords {
			if u == password {
				passwordCorrect:=true
				break
			}
		}
	}

	fmt.print("You are singed in.")
}
