package bank

import (
	"errors"
	"fmt"
)

func (b *Bank) Deposit(amount float64, password string) error {

	if password == b.Password {
		b.Balance += amount
		fmt.Printf("Deposited %f to account number %d. New balance is %f\n", amount, b.AccountNumber, b.Balance)
	} else {
		return errors.New("Invalid password. Deposit failed.")
	}
	return nil
}
func (b *Bank) Withdraw(amount float64, password string) error {
	if password != b.Password {
		return errors.New("invalid password: withdrawal failed")
	}
	if amount > b.Balance {
		return errors.New("insufficient balance: withdrawal failed")
	}
	b.Balance -= amount
	fmt.Printf("Withdrew %.2f from account %d. New balance: %.2f\n", amount, b.AccountNumber, b.Balance)
	return nil
}

func (b *Bank) Transfer(amount float64, recipient *Bank, password string) error {
	if password == b.Password {
		if amount <= b.Balance {
			if recipient != nil && recipient.AccountNumber != 0 && Banks[recipient.AccountNumber] != nil {
				b.Balance -= amount
				recipient.Balance += amount
				fmt.Printf("Transferred %f from account number %d to account number %d. New balance is %f\n", amount, b.AccountNumber, recipient.AccountNumber, b.Balance)
			} else {
				return errors.New("Invalid recipient account. Transfer failed.")
			}
		} else if amount > b.Balance {
			return errors.New("Insufficient balance. Transfer failed.")
		}
	} else {
		return errors.New("Invalid password. Transfer failed.")
	}
	return nil
}
