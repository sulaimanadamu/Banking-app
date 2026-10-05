package profile

import (
	"bankApp/util"
	"errors"
	"fmt"
	"strconv"
)

var name string
var email string
var userName string
var pin int64

func CreateUser(name, userName string, pin int64) (string, error) {
	// create file by user name in /UserData/
	src := "./UserData/"
	format := ".txt"
	filePath := src + userName + format
	util.WriteTo(filePath, util.IntToString(pin))
	util.WriteTo(filePath, name)
	util.WriteTo(filePath, "0.00")
	return filePath, nil
}

// return username, this is then used as file search name.
func Login(userName string, pin int64) (string, error) {
	src := "./UserData/"
	format := ".txt"
	filePath := src + userName + format

	_, err := authorization(userName, pin)
	// using error instead of return value, because its more descriptive.
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Username, password or both may not be correct try again.")
	} else {
		fmt.Printf("Successfully logged in user %s", userName)
		return filePath, nil
	}
}

func authorization(userName string, pin int64) (bool, error) {
	src := "../UserData/"
	format := ".txt"
	filename := src + userName + format

	// This check if the user exist by searching for his/her documents
	if util.FileExist(filename) {
		savedPin, err := getSavedPin(filename)
		if err != nil {
			return false, err
		} else {
			if savedPin == pin {
				return true, nil
			} else {
				return false, errors.New("incorrect pin")
			}
		}
	}
	return false, errors.New("authorization failed.")
}

// checks if user exist then does the pin match current record
func getSavedPin(filename string) (int64, error) {
	// pin is saved as first line in every user file.
	pinString := util.GetTextLineValue(filename, 1)

	savedPin, err := strconv.ParseInt(pinString, 10, 64)
	if err != nil {
		return 0, errors.New("pin can only contain numbers.")
	} else {
		return savedPin, nil
	}
}
