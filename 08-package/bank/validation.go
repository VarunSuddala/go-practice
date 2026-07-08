package bank

import (
	"errors"
	"fmt"
)

func (b *Bank) ChangePassword(oldPassword, newPassword string) error {
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

func (b *Bank) ChangeEmail(newEmail string, password string) error {
	if password == b.Password {
		b.Email = newEmail
		fmt.Println("Email changed successfully.")
	} else {
		return errors.New("Invalid password. Email change failed.")
	}
	return nil
}
