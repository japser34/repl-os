package startup

import (
	"os"
	"encoding/json"
	"errors"
	"path/filepath"
	"golang.org/x/crypto/bcrypt"
)

const fileName = "data/users.json"

type User struct {
	Username string
	Password string
}

func GetUserData() ([]User, error) {
	file, err := os.Open(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return writeDefaultUserData()
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}


func writeDefaultUserData() (userData []User, err error) {
	users := []User{
		{"Jasper", "HelloWorld"},
	}

	for i, user := range users {
		users[i].Password, err = hashPassword(user.Password)
		if err != nil {
			return nil, err
		}
	}

	data, err := json.MarshalIndent(users, "" , "  ")
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(filepath.Dir(fileName), 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	_, err = file.Write(data)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
