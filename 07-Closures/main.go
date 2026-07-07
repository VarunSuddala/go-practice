// Beginner
// Counter
// ID Generator
// Login Attempts
// Intermediate
// Wallet
// Bank Account
// Multiplier Factory
// Discount Factory
// Advanced
// Fibonacci Generator
// Even Generator
// Prime Generator
// Random Password Generator

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")
	counterFunc := counter()
	fmt.Println(counterFunc())
	fmt.Println(counterFunc())
	idgenerator := idgenerator()
	fmt.Println(idgenerator())
	fmt.Println(idgenerator())
	loginAttempts := loginAttempts()
	fmt.Println(loginAttempts())
	fmt.Println(loginAttempts())
	fmt.Println(loginAttempts())
	fmt.Println(loginAttempts())
	walletDeposit, walletBalance := wallet()
	fmt.Println(walletDeposit(100.0))
	fmt.Println(walletDeposit(100.0))
	fmt.Println(walletBalance())
	bankDeposit, bankWithdraw, bankBalance := bankAccount()
	fmt.Println(bankDeposit(500.0))
	fmt.Println(bankWithdraw(200.0))
	fmt.Println(bankBalance())
	fmt.Println(bankWithdraw(400.0))
	fmt.Println(bankBalance())
	multiplierBy2 := multiplierFactory(2)
	fmt.Println(multiplierBy2(5))
	fmt.Println(multiplierBy2(10))
	discount10 := discountFactory(0.1)
	fmt.Println(discount10(100.0))
	fmt.Println(discount10(200.0))
	evenGen := evenGenerator()
	fmt.Println(evenGen())
	fmt.Println(evenGen())
	fmt.Println(evenGen())
	fibGen := fibonacciGenerator()
	fmt.Println(fibGen())
	fmt.Println(fibGen())
	fmt.Println(fibGen())
	primeGen := primeGenerator()
	fmt.Println(primeGen())
	fmt.Println(primeGen())
	fmt.Println(primeGen())
	passwordGen := randomPasswordGenerator(10)
	fmt.Println(passwordGen())
	fmt.Println(passwordGen())
}

func counter() func() int {
	coun := 0
	inner := func() int {
		coun += 1
		return coun
	}
	return inner
}

func idgenerator() func() int {
	id := 1000
	inner := func() int {
		id++
		return id
	}
	return inner
}

func loginAttempts() func() int {
	attempts := 0
	inner := func() int {
		if attempts < 3 {
			attempts++
			return attempts
		} else {
			fmt.Println("Maximum login attempts reached.")
			return attempts
		}
	}
	return inner
}
func wallet() (func(amount float64) float64, func() float64) {
	balance := 0.0
	deposit := func(amount float64) float64 {
		balance += amount
		return balance
	}
	getBalance := func() float64 {
		return balance
	}
	return deposit, getBalance
}

func bankAccount() (func(amount float64) float64, func(amount float64) float64, func() float64) {
	balance := 0.0
	deposit := func(amount float64) float64 { balance += amount; return balance }
	withdraw := func(amount float64) float64 {
		if amount > balance {
			fmt.Println("Insufficient funds.")
			return balance
		}
		balance -= amount
		return balance
	}
	getBalance := func() float64 { return balance }
	return deposit, withdraw, getBalance
}

func multiplierFactory(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func discountFactory(discount float64) func(float64) float64 {
	return func(price float64) float64 {
		return price * (1 - discount)
	}
}

func fibonacciGenerator() func() int {
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func evenGenerator() func() int {
	n := 0
	return func() int {
		n += 2
		return n
	}
}

func primeGenerator() func() int {
	n := 1
	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	return func() int {
		for {
			n++
			if isPrime(n) {
				return n
			}
		}
	}
}

func randomPasswordGenerator(length int) func() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%+"
	return func() string {
		password := make([]byte, length)
		for i := range password {
			password[i] = charset[rand.Intn(len(charset))]
		}
		return string(password)
	}
}
