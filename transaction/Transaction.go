package transaction

import (
	"fmt"
	"os"
	"strconv"
)

var accountBalance float64 = 10000.00
var accountBalanceFile string = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func getBalanceFromFile() (balance float64) {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ = strconv.ParseFloat(balanceText, 64)
	return
}

func Transaction() {
	fmt.Println("Welcome to soul Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	for i := 1; i > 0; i++ {
		fmt.Print("Enter an input: ")
		choice := getInput()

		switch choice {
		case 1:
			fmt.Println("your account balance is:", accountBalance)
		case 2:
			fmt.Print("Enter a Deposit: ")
			credit := float64(getInput())
			if credit > 0 {
				accountBalance += credit
				fmt.Println("The new balance is:", accountBalance)
				writeBalanceToFile(accountBalance)
			} else {
				fmt.Println("No zero or negative deposit")
			}
		case 3:
			fmt.Print("Enter withdrawal amount: ")
			debit := float64(getInput())
			if debit > accountBalance {
				fmt.Println("can't withdraw this much!, your current balance is: ", accountBalance)
			} else {
				accountBalance -= float64(debit)
				fmt.Println("Your new balance:", accountBalance)
			}
		case 4:
			fmt.Println("EXIT")
			fmt.Println("Thanks for banking with us !")
			return
		default:
			fmt.Println("Invalid input")
		}
	}

}

func getInput() (input int) {
	fmt.Scan(&input)
	return
}
