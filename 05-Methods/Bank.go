package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Bank struct {
	Name          string
	Email         string
	AccountNumber int64
	Balance       float64
	Address       string
	Type          string
	Password      string
}
var banks = make(map[int64] *Bank)

func (b Bank) display(password string) error {
	if password == b.Password {
		fmt.Printf("name : %s \nEmail : %s\nAccountNumber : %d\n Balance :%f\nType : %s", b.Name, b.Email, b.AccountNumber, b.Balance, b.Type)
	} else {
		return errors.New("Invalid password. Display failed.")
	}
	return nil
}

func (b *Bank) deposit(amount float64, password string) error {

	if password == b.Password {
		b.Balance += amount
		fmt.Printf("Deposited %f to account number %d. New balance is %f\n", amount, b.AccountNumber, b.Balance)
	} else {
		return errors.New("Invalid password. Deposit failed.")
	}
	return nil
}
func (b *Bank) withdraw(amount float64, password string) error {
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

func (b *Bank) transfer(amount float64, recipient *Bank, password string) error {
	if password == b.Password {
		if amount <= b.Balance {
			if recipient != nil && recipient.AccountNumber != 0 && banks[recipient.AccountNumber] != nil {
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

func (b *Bank) changePassword(oldPassword, newPassword string) error {
    if oldPassword != b.Password {
        return errors.New("invalid old password")
    }
    if len(newPassword) != 4 {
        return errors.New("new password must be 4 digits")
    }
    b.Password = newPassword
    fmt.Println("Password changed successfully.")
    return nil
}

func (b *Bank) changeEmail(newEmail string, password string) error {
	if password == b.Password {
		b.Email = newEmail
		fmt.Println("Email changed successfully.")
	} else {
		return errors.New("Invalid password. Email change failed.")
	}
	return nil
}

func makeAccount(name string, email string, balance float64, address string, accountType string, password string) (*Bank, error) {
	
	b := Bank{}
	var accountNumber int64
	for {
		accountNumber = rand.Int63n(9_000_000_000_000) + 1_000_000_000_000
        if _, exists := banks[accountNumber]; !exists {
            break
        }
		
	}
	
	if name == "" || email == "" || address == "" || password == "" {
		return nil, errors.New("Invalid input. Account creation failed.")
	}

	if len(password) != 4 {
		return nil, errors.New("Password must be 4 digits. Account creation failed.")
	}

	if balance < 0 {
		fmt.Println("Invalid balance. Account creation failed.")
		return nil, errors.New("Invalid balance. Account creation failed.")
	}

	if accountType != "savings" && accountType != "current" {
		return nil, errors.New("Invalid account type. Account creation failed.")
	}
	b.Name = name
	b.Email = email
	b.AccountNumber = accountNumber
	b.Balance = balance
	b.Address = address
	b.Type = accountType
	b.Password = password
	fmt.Println("Account created successfully.")
	
	return &b, nil
}
