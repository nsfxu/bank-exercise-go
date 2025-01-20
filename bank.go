package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(balanceFileName)

	if err != nil {
		fmt.Println("--------------")
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 at ", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: $%.2f\n", accountBalance)
		case 2:
			fmt.Println("What is the amount that you want to deposit?")
			fmt.Print("Amount: ")

			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount

			fmt.Printf("Balance updated! Now your current balance is: $%.2f", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFileName)
		case 3:
			fmt.Println("What is the amount that you want to withdraw?")
			fmt.Print("Amount: ")

			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount

			fmt.Printf("Balance updated! Now your current balance is: $%.2f", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFileName)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again")
			continue
		}

	}

}
