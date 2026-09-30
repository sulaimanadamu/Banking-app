package profile

import (
	"bankApp/util"
	"errors"
	"os"
	"strconv"
)

var name string
var email string
var userName string
var pin int64

func createUser(name, userName string, pin int) {
	// create file by user name in /UserData/
	fileName := "../UserData/" + userName + ".txt"
	_, err := os.Stat(fileName)
	if err != nil {

	} else {
		// no error
		// append per line pin, name, email, password to the file.
		util.WriteTo(userName, strconv.Itoa(pin))
		util.WriteTo(userName, "name: "+name)

	}

}

func authorization(userName string, pin int64) (bool, error) {
	src := "../UserData/"
	// check if the user id matches password
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
