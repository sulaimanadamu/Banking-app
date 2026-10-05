package main

import (
	"bankApp/profile"
	"bankApp/transaction"
	"bankApp/util"
	"fmt"
	"strconv"
)

func main() {
	//data file path

	// want to login or sign up?
	fmt.Println("Welcome to Codex Bank")
	fmt.Println("Select option for details? ")
	fmt.Println("1. Login")
	fmt.Println("2. Signup")
	input := util.GetInput()
	option, convertErr := util.StringToInt(input)

	if convertErr != nil {
		fmt.Printf("this should run!, possible err is %v", convertErr)
	} else {
		switch option {
		case 1:
			// login
			fmt.Print("Enter a user name: ")
			userName := util.GetInput()
			fmt.Print("Enter a pin: ")
			pinString := util.GetInput()
			pin, _ := util.StringToInt(pinString)
			filePath, loginErr := profile.Login(userName, pin)

			if loginErr != nil {
				fmt.Println("can't login!")
			} else {
				transaction.Transaction(filePath)
			}

		case 2:
			fmt.Print("Enter a user name: ")
			userName := util.GetInput()

			fmt.Print("Enter a pin: ")
			pin, _ := strconv.ParseInt(util.GetInput(), 10, 64)

			fmt.Print("Enter a your name: ")
			name := util.GetInput()

			filePath, err := profile.CreateUser(name, userName, pin)
			if err != nil {
				fmt.Println("user doesn't exist!")
			} else {
				transaction.Transaction(filePath)
			}

		default:
			fmt.Println("enter only a valid input")

		}
	}

}
