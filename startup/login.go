package startup

import (
	"fmt"
)

func Login() {
	usernames := []string{"Jasper"}


	var username string
	var usernameCorrect := false
	fmt.Println("Please login")
	for !usernameCorrect { 
		fmt.Print("Username: ")
		fmt.Scan(&username)
		for u := range usernames {
			if u == username {
				usernameCorrect:=true
				break
			}
		}
	}
}
