package main

import (
	"fmt"

	"github.com/VarunSuddala/go-learning/08-package/bank"
)

func main() {
	for {
		var choice int
		fmt.Println("1. Create Account")
		fmt.Println("2. Display Account Details")
		fmt.Println("3. Deposit")
		fmt.Println("4. Withdraw")
		fmt.Println("5. Transfer")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, email, address, accountType, password string
			var balance float64
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter email: ")
			fmt.Scan(&email)
			fmt.Print("Enter initial balance: ")
			fmt.Scan(&balance)
			fmt.Print("Enter address: ")
			fmt.Scan(&address)
			fmt.Print("Enter account type (savings/current): ")
			fmt.Scan(&accountType)
			fmt.Print("Set a 4-digit password: ")
			fmt.Scan(&password)
			account, err := bank.MakeAccount(name, email, balance, address, accountType, password)
			if err != nil {
				fmt.Println("Account creation failed:", err)
			} else {
				bank.Banks[account.AccountNumber] = account
				fmt.Println("Account created successfully. Your account number is:", account.AccountNumber)
			}

		case 2:
			var accountNumber int64
			var password string
			fmt.Print("Enter account number: ")
			fmt.Scan(&accountNumber)
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			Bank, exists := bank.Banks[accountNumber]
			if exists {
				Bank.Display(password)
			} else {
				fmt.Println("Account not found.")
			}

		case 3:
			var accountNumber int64
			var amount float64
			var password string
			fmt.Print("Enter account number: ")
			fmt.Scan(&accountNumber)
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			Bank, exists := bank.Banks[accountNumber]
			if exists {
				Bank.Deposit(amount, password)
			} else {
				fmt.Println("Account not found.")
			}

		case 4:
			var accountNumber int64
			var amount float64
			var password string
			fmt.Print("Enter account number: ")
			fmt.Scan(&accountNumber)
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			Bank, exists := bank.Banks[accountNumber]
			if exists {
				Bank.Withdraw(amount, password)
			} else {
				fmt.Println("Account not found.")
			}

		case 5:
			var senderAccountNumber, recipientAccountNumber int64
			var amount float64
			var password string
			fmt.Print("Enter your account number: ")
			fmt.Scan(&senderAccountNumber)
			fmt.Print("Enter recipient's account number: ")
			fmt.Scan(&recipientAccountNumber)
			fmt.Print("Enter amount to transfer: ")
			fmt.Scan(&amount)
			fmt.Print("Enter your password: ")
			fmt.Scan(&password)
			senderBank, senderExists := bank.Banks[senderAccountNumber]
			recipientBank, recipientExists := bank.Banks[recipientAccountNumber]

			if senderExists && recipientExists {
				err := senderBank.Transfer(amount, recipientBank, password)
				if err != nil {
					fmt.Println("Transfer failed:", err)
				} else {
					fmt.Println("Transfer successful.")
					fmt.Printf("New balance for account number %d: %f\n", senderAccountNumber, senderBank.Balance)
				}
			} else {
				fmt.Println("One or both accounts not found.")
			}

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}
