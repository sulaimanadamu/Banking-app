package transaction

import (
	"bankApp/util"
	"errors"
	"fmt"
	"strconv"
)

func Transaction(filePath string) error {
	lastLine := util.GetTextLineValue(filePath, -1)

	accountBalance, _ := util.StringToFloat(lastLine)

	fmt.Printf("Welcome to soul Bank! %s \n", util.GetTextLineValue(filePath, 2))
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	for i := 1; i > 0; i++ {
		fmt.Print("Enter an input: ")
		userInput := util.GetInput()
		fmt.Println(userInput)
		choice, err := util.StringToInt(userInput)

		if err != nil {
			nwErr := errors.New("enter an integer value")
			return nwErr
		} else {
			switch choice {
			case 1:
				fmt.Println("your account balance is: ", util.GetTextLineValue(filePath, -1))
			case 2:
				fmt.Print("Enter a Deposit: ")
				credit, err := strconv.ParseFloat(util.GetInput(), 64)
				if err != nil {
					return errors.New("enter your value in decimal and make sure they are all numbers.")
				} else {
					if credit > 0 {
						accountBalance += credit
						fmt.Println("The new balance is: ", accountBalance)
						formattedLine := util.FloatToString(accountBalance)
						util.WriteTo(filePath, formattedLine)

					} else {
						fmt.Println("No zero or negative deposit")
					}
				}

			case 3:
				fmt.Print("Enter withdrawal amount: ")
				debit, err := util.StringToFloat(util.GetInput())

				if err != nil {

				} else {

					if debit > accountBalance {
						fmt.Println("can't withdraw this much!, your current balance is: ", accountBalance)
					} else {

						accountBalance -= float64(debit)
						fmt.Println("Your new balance:", accountBalance)
						formattedLine := util.FloatToString(accountBalance)
						util.WriteTo(filePath, formattedLine)

					}
				}
			case 4:
				fmt.Println("EXIT")
				fmt.Println("Thanks for banking with us !")
				return nil
			default:
				fmt.Println("Invalid input")
			}
		}
	}
	return nil
}
