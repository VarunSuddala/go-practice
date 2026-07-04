package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func inventory() map[string]Product {

	inventory := make(map[string]Product)

	inventory["P001"] = Product{Name: "Phone", Price: 699.99, Quantity: 10}
	inventory["P002"] = Product{Name: "Laptop", Price: 1299.99, Quantity: 5}
	inventory["P003"] = Product{Name: "Tablet", Price: 399.99, Quantity: 8}

	return inventory
}

func addProduct(inventory map[string]Product, productID string, product Product) {
	inventory[productID] = product
}

func removeProduct(inventory map[string]Product, productID string) {
	delete(inventory, productID)

}

func updateProductQuantity(inventory map[string]Product, productID string, newQuantity int) {
	if product, exists := inventory[productID]; exists {
		product.Quantity = newQuantity
		inventory[productID] = product
	}
}

func retrieveProduct(inventory map[string]Product, productID string) (Product, bool) {
	product, exists := inventory[productID]
	return product, exists
}
func incrementProductQuantity(inventory map[string]Product, productID string, increment int) {
	if product, exists := inventory[productID]; exists {
		product.Quantity += increment
		inventory[productID] = product
	}
}

func dropProductQuantity(inventory map[string]Product, productID string, decrement int) {
	if product, exists := inventory[productID]; exists {
		if product.Quantity >= decrement {
			product.Quantity -= decrement
			inventory[productID] = product
		}
	}
}

func updatePrice(inventory map[string]Product, productID string, newPrice float64) {
	if product, exists := inventory[productID]; exists {
		product.Price = newPrice
		inventory[productID] = product
	}
}
func totalInventoryValue(inventory map[string]Product) float64 {
	totalValue := 0.0
	for _, product := range inventory {
		totalValue += product.Price * float64(product.Quantity)
	}
	return totalValue
}

func printInventory(inventory map[string]Product) {
	fmt.Println("Inventory:")
	for productID, product := range inventory {
		fmt.Printf("Product ID: %s, Name: %s, Price: %.2f, Quantity: %d\n", productID, product.Name, product.Price, product.Quantity)
	}
}
