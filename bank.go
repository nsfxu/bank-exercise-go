package main

import "fmt"

func main() {

	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Your balance is: $%.2f", accountBalance)
	} else if choice == 2 {
		fmt.Println("What is the amount that you want to deposit?")
		fmt.Print("Amount: ")

		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		accountBalance += depositAmount

		fmt.Printf("Balance updated! Now your current balance is: $%.2f", accountBalance)
	} else if choice == 3 {
		fmt.Println("What is the amount that you want to withdraw?")
		fmt.Print("Amount: ")

		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have.")
			return
		}

		accountBalance -= withdrawAmount

		fmt.Printf("Balance updated! Now your current balance is: $%.2f", accountBalance)
	} else if choice == 4 {
		fmt.Println("Goodbye!")
	}
}
