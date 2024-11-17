package startup

import (
	"os"
	"encoding/json"
	"errors"
)

const fileName = "userData.json"

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

	data, err := json.MarshalIndent(users, "" , "  ")
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
