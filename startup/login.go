
package startup

import (
	"fmt"
)

func Login() {
	usernames := []string{"Jasper"}
	passwords := []string{"HelloWorld"}

	fmt.Println("Please login")
	var usernameCorrect bool
	for !usernameCorrect {
		var username string
		fmt.Print("Username: ")
		fmt.Scan(&username)
		for _, u := range usernames {
			if u == username {
				usernameCorrect = true
				break
			}
		}
		if !usernameCorrect {
			fmt.Println("Invalid username. Please try again.")
		}
	}

	var passwordCorrect bool
	for !passwordCorrect {
		var password string
		fmt.Print("Password: ")
		fmt.Scan(&password)
		for _, p := range passwords {
			if p == password {
				passwordCorrect = true
				break
			}
		}
		if !passwordCorrect {
			fmt.Println("Invalid password. Please try again.")
		}
	}

	fmt.Println("You are signed in.")
}

