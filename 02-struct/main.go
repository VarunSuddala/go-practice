// Create a Book struct with Title, Author, and Price.
// Create three books and print their details.
// Increase the price of one book.
// Create a nested Address inside an Employee struct.
// Write a function that takes *Employee and updates the employee's city.
// Add JSON tags to a User struct and marshal it to JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Price  float64
}
type Address struct {
	City  string
	State string
}
type Employee struct {
	Name    string
	Age     int
	Email   string
	Address Address
}
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("Structs in Golang")
	book1 := createBook("The Go Programming Language", "Alan A. A. Donovan", 39.99)
	book2 := createBook("The Go Standard Library", "Alan A. A. Donovan", 29.99)
	book3 := createBook("Learning Go Programming", "Rustan Leino", 24.99)

	fmt.Printf("Book 1: %s by %s - $%.2f\n", book1.Title, book1.Author, book1.Price)
	fmt.Printf("Book 2: %s by %s - $%.2f\n", book2.Title, book2.Author, book2.Price)
	fmt.Printf("Book 3: %s by %s - $%.2f\n", book3.Title, book3.Author, book3.Price)

	increasePrice(&book2, 50.00)
	fmt.Println("after upadate ", book2)
	addr := Address{
		City:  "warangal",
		State: "Telangana",
	}

	emp := createEmployee("varun", 22, "varun@gmail.com", addr)
	fmt.Println(emp)

	updateCity(&emp, "hanamkonda")
	fmt.Println(&emp)

	user := User{
		Id:    1,
		Name:  "varun",
		Email: "varun@gmail.com",
	}

	jsonParsing(user)

}

func createBook(title string, author string, price float64) Book {
	book := Book{
		Title:  title,
		Author: author,
		Price:  price,
	}
	return book
}

func increasePrice(book *Book, price float64) {
	book.Price += price
}

func createEmployee(name string,
	age int,
	email string,
	address Address) Employee {
	emp := Employee{
		Name:    name,
		Age:     age,
		Email:   email,
		Address: address,
	}
	return emp
}

func updateCity(emp *Employee, city string) {
	emp.Address.City = city
}
func jsonParsing(user User) {

	data, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

}
