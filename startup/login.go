package startup

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)


func Login() {
	users, err := GetUserData()
	if err != nil {
		fmt.Println("Error getting user data:", err)
		return
	}

	fmt.Println("Please login")
	var username, password string
	for {
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		for _, user := range users {
			if user.Username == username {
				err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
				if err == nil {
					fmt.Println("Login successful.")
					return
				}
			}
		}
		fmt.Println("Invalid username or password.")
	}
}

