package bank

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

var Banks = make(map[int64]*Bank)

func (b Bank) Display(password string) error {
	if password == b.Password {
		fmt.Printf("name : %s \nEmail : %s\nAccountNumber : %d\n Balance :%f\nType : %s", b.Name, b.Email, b.AccountNumber, b.Balance, b.Type)
	} else {
		return errors.New("Invalid password. Display failed.")
	}
	return nil
}

func MakeAccount(name string, email string, balance float64, address string, accountType string, password string) (*Bank, error) {

	b := Bank{}
	var accountNumber int64
	for {
		accountNumber = rand.Int63n(9_000_000_000_000) + 1_000_000_000_000
		if _, exists := Banks[accountNumber]; !exists {
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
